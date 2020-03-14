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
	conn, err := redis.Dial("tcp", ":6379")
	testutil.AssertNil(t, err)
	queue := &Queue{conn: conn}
	err = queue.Produce(msg)
	testutil.AssertNil(t, err)
}
