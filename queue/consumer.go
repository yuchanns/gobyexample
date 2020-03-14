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
	doingName := fmt.Sprintf("%s.doing", c.name)

	go func() {
		fmt.Println("start loop...")
		for {
			// BROPLPUSH sourceQueueName destQueueName timeout
			if r, err := c.client.Do("BRPOPLPUSH", prepareName, doingName, 10); err == nil {
				if rUint8s, ok := r.([]uint8); ok {
					msg := &Message{}
					if err := jsoniter.Unmarshal(rUint8s, msg); err == nil {
						if c.Consume(msg) {
							// LREM queueName numbers msg
							if _, err := c.client.Do("LREM", doingName, 1, r); err != nil {
								fmt.Println("failed to lrem", err)
							}
						}
					} else {
						fmt.Println("failed to pop msg", err)
					}
				} else {
					fmt.Println("failed to convert reply to []uint8")
				}
			} else {
				fmt.Println(err)
			}
			// ack
			for {
				if reply, err := c.client.Do("RPOPLPUSH", doingName, prepareName); err != nil {
					fmt.Println("error happens", err)
				} else {
					if reply == nil {
						break
					} else {
						fmt.Printf("got undo msg in the queue %s\n", doingName)
					}
				}
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
