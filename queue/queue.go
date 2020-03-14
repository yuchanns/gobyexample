package main

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// IMessage is the interface which Queue accept
type IMessage interface {
	Consume() error
	GetChannel() string
	Marshal() ([]byte, error)
	Unmarshal([]uint8) (IMessage, error)
}

// Queue can produce and consume messages
// TODO: A Queue should have context to cancel the consumer process
type Queue struct {
	conn redis.Conn
}

func (q *Queue) lrem(queue string, reply interface{}) {
	// redis tip: LREM queueName numbers msg
	if _, err := q.conn.Do("LREM", queue, 1, reply); err != nil {
		fmt.Println("failed to lrem", err)
	}
}

func (q *Queue) rpoplpush(imsg IMessage, sourceQueue, destQueue string, block bool) (interface{}, IMessage, error) {
	var r interface{}
	var err error
	if block {
		// redis tip: BROPLPUSH sourceQueueName destQueueName timeout
		r, err = q.conn.Do("BRPOPLPUSH", sourceQueue, destQueue, 10)
	} else {
		// redis tip: ROPLPUSH sourceQueueName destQueueName
		r, err = q.conn.Do("RPOPLPUSH", sourceQueue, destQueue)
	}
	if err != nil {
		return nil, nil, err
	}
	// return all nil when timeout and read nothing
	if r == nil {
		return nil, nil, nil
	}
	rUint8, ok := r.([]uint8)
	if !ok {
		return nil, nil, errors.New("cannot assert reply as type []uint8")
	}

	if msg, err := imsg.Unmarshal(rUint8); err != nil {
		return nil, nil, err
	} else if _, ok := msg.(IMessage); ok {
		return r, msg, nil
	} else {
		return nil, nil, errors.New("cannot assert msg as interface IMessage")
	}
}

func (q *Queue) ack(imsg IMessage, sourceQueue, destQueue string) {
	for {
		reply, _, err := q.rpoplpush(imsg, sourceQueue, destQueue, false)
		if err != nil {
			fmt.Println("ack failed", err)
			break
		}
		if reply == nil {
			break
		} else {
			fmt.Printf("got undo msg in the queue %s\n", sourceQueue)
		}
	}
}

// InitConsumer will create goroutine to consume the msg implementing IMessage
// the number decides how much goroutine to do consuming
func (q *Queue) InitConsumer(msg IMessage, number int) {
	prepareQueue := fmt.Sprintf("%s.prepare", msg.GetChannel())
	doingQueue := fmt.Sprintf("%s.doing", msg.GetChannel())

	if number <= 0 {
		number = 1
	}

	for i := 0; i < number; i++ {
		go func() {
			for {
				reply, msg, err := q.rpoplpush(msg, prepareQueue, doingQueue, true)
				if err != nil {
					fmt.Println("failed to pop msg", err)
					continue
				}
				if msg == nil {
					continue
				}
				if err := msg.Consume(); err == nil {
					q.lrem(doingQueue, reply)
				}
				q.ack(msg, doingQueue, prepareQueue)
			}
		}()
	}
	fmt.Println("consumer initialed")
}

// Produce will produce msg implementing IMessage into the queue to be consumed
func (q *Queue) Produce(msg IMessage) error {
	prepareQueue := fmt.Sprintf("%s.prepare", msg.GetChannel())
	if msgJson, err := msg.Marshal(); err != nil {
		return err
	} else {
		_, err := q.conn.Do("LPUSH", prepareQueue, msgJson)
		fmt.Println("produced", string(msgJson))

		return err
	}
}
