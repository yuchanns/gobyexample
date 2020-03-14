package main

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	pool := &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", ":6379")
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	queue := &Queue{pool: pool}

	msg := &Message{
		name: "demoQueue",
	}

	// make a main context
	ctx := context.Background()
	// pass the main context to queue
	cancelFunc := queue.InitReceiver(ctx, msg, 10)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT)

	for {
		switch <-quit {
		case syscall.SIGINT:
			cancelFunc()
			return
		}
	}
}
