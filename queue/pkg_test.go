package main

import (
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"testing"
	"time"
)

func TestQueue_Produce(t *testing.T) {
	msg := &Message{name: "demoQueue", Content: map[string]string{
		"order_no": strconv.FormatInt(time.Now().Unix(), 10),
	}}
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
	for {
		err := queue.Delivery(msg)
		testutil.AssertNil(t, err)
		time.Sleep(time.Millisecond)
	}
}
