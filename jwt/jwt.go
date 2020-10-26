package jwt

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

var (
	pool         *redis.Pool
	hmacSecret   = []byte("aabbccdd")
	t            = 86400 * time.Second
	redisAddress = ":6379"
	redisPwd     = ""
	redisPrefix  = "mp"
)

func init() {
	pool = &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", redisAddress, redis.DialPassword(redisPwd))
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

type Jwt struct {
	pool  *redis.Pool
	token *jwt.Token
	t     time.Duration
}

func NewJwt() *Jwt {
	token := jwt.New(jwt.SigningMethodHS256)
	return &Jwt{
		pool:  pool,
		token: token,
		t:     t,
	}
}

func (j *Jwt) Generate(id int, name string) (string, error) {
	j.token.Claims = jwt.MapClaims{
		"id":   id,
		"name": name,
	}

	tokenString, err := j.token.SignedString(hmacSecret)
	if err == nil {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		conn, err := j.pool.GetContext(ctx)
		if err != nil {
			return "", err
		}
		defer conn.Close()

		key, err := j.buildKey(j.token.Claims)

		if err != nil {
			return "", err
		}

		if err := j.setEx(conn, key, tokenString); err != nil {
			return "", err
		}
	}

	return tokenString, err
}

func (j *Jwt) Validate(tokenString string) error {
	token, err := j.parseToken(tokenString)

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("token validate failed")
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := j.pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	key, err := j.buildKey(token.Claims)
	if err != nil {
		return err
	}

	// to check if exist and expired
	if _, err := j.checkIfExpired(conn, key); err != nil {
		return err
	}

	// extend expire date
	return j.setEx(conn, key, tokenString)
}

func (j Jwt) Invalidate(tokenString string) error {
	token, err := j.parseToken(tokenString)

	if err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := j.pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	return j.delKey(conn, token.Claims)
}

func (j *Jwt) parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSecret, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (j *Jwt) buildKey(jwtCaims jwt.Claims) (string, error) {
	claims, ok := jwtCaims.(jwt.MapClaims)
	if !ok {
		return "", errors.Errorf("Unexpected claims type: %T, value: %v", claims, claims)
	}
	id, ok := claims["id"].(int)
	if !ok {
		idFloat64, ok := claims["id"].(float64)
		if !ok {
			return "", errors.Errorf("Unexpected claims' id type: %T, value: %v", claims["id"], claims["name"])
		}
		id = int(idFloat64)
	}
	name, ok := claims["name"].(string)
	if !ok {
		return "", errors.Errorf("Unexpected claims' name type: %T, value: %v", claims["name"], claims["name"])
	}

	return strings.Join(
		[]string{redisPrefix, strconv.Itoa(id), name},
		"_",
	), nil
}

func (j *Jwt) setEx(conn redis.Conn, key, val string) error {
	_, err := conn.Do("SETEX", key, j.t.Seconds(), val)

	return err
}

func (j *Jwt) checkIfExpired(conn redis.Conn, key string) (time.Duration, error) {
	ttl, err := redis.Int64(conn.Do("TTL", key))
	if err != nil {
		return 0, err
	}

	ttlDuration := time.Duration(ttl)

	if ttlDuration <= 0 {
		return 0, errors.New("the token has been expired")
	}

	return ttlDuration, nil
}

func (j *Jwt) delKey(conn redis.Conn, jwtCaims jwt.Claims) error {
	key, err := j.buildKey(jwtCaims)
	if err != nil {
		return err
	}
	_, err = conn.Do("DEL", key)

	return err
}
