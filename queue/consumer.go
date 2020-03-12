package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	jsoniter "github.com/json-iterator/go"
	"os"
	"os/signal"
	"syscall"
)

type IConsumer interface {
	Consume(message *Message) bool
	Ack()
	GetChannel() string
}

type consumer struct {
	consumer func(message *Message) bool
	name     string
	client   redis.Conn
}

func (c *consumer) Consume(message *Message) bool {
	return c.consumer(message)
}

func (c *consumer) GetChannel() string {
	return c.name
}

func (c *consumer) Ack() {
	prepareName := fmt.Sprintf("%s.prepare", c.name)

	go func() {
		fmt.Println("start loop...")
		for {
			if r, err := redis.ByteSlices(c.client.Do("BRPOP", prepareName, 10)); err == nil {
				msg := &Message{}
				if err := jsoniter.Unmarshal(r[1], msg); err == nil {
					c.Consume(msg)
				} else {
					fmt.Println("failed to pop msg", err)
				}
			} else {
				fmt.Println(err)
			}
		}
	}()
	fmt.Println("consumer start...")

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT)

	for {
		switch <-quit {
		case syscall.SIGINT:
			os.Exit(0)
		default:
			return
		}
	}
}

func NewConsumer(f func(message *Message) bool, name string, client redis.Conn) *consumer {
	return &consumer{
		consumer: f,
		name:     name,
		client:   client,
	}
}
