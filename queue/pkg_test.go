package main

import (
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/gomodule/redigo/redis"
	"testing"
	"time"
)

func TestProducer_Produce(t *testing.T) {
	c, err := redis.Dial("tcp", ":6379")
	testutil.AssertNil(t, err)
	p := NewProducer("test", c)
	p.Produce(&Message{
		Content: []string{
			"hello",
			"world",
		},
		Id: time.Now().Unix(),
	})
}
