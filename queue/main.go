package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	m := NewConsumer(func(msg *Message) bool {
		fmt.Println("Got the message content is", msg.Content)
		return true
	}, "test", c)
	DoConsume(m)
}

func DoConsume(c IConsumer) {
	c.Ack()
}
