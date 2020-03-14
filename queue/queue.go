package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// IMessage is the interface which Queue accept
type IMessage interface {
	Resolve() error
	GetChannel() string
	Marshal() ([]byte, error)
	Unmarshal([]uint8) (IMessage, error)
}

// Queue can produce and consume messages
type Queue struct {
	pool *redis.Pool
}

func (q *Queue) lrem(queue string, reply interface{}) {
	// redis tip: LREM queueName numbers msg
	conn := q.pool.Get()
	defer conn.Close()
	if _, err := conn.Do("LREM", queue, 1, reply); err != nil {
		fmt.Println("failed to lrem", err)
	}
}

func (q *Queue) rpoplpush(imsg IMessage, sourceQueue, destQueue string, block bool) (interface{}, IMessage, error) {
	var r interface{}
	var err error
	conn := q.pool.Get()
	defer conn.Close()
	if block {
		// redis tip: BROPLPUSH sourceQueueName destQueueName timeout
		r, err = conn.Do("BRPOPLPUSH", sourceQueue, destQueue, 1)
	} else {
		// redis tip: ROPLPUSH sourceQueueName destQueueName
		r, err = conn.Do("RPOPLPUSH", sourceQueue, destQueue)
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

// InitReceiver will create goroutine to consume the msg implementing IMessage
// the number decides how much goroutine to do consuming
func (q *Queue) InitReceiver(ctx context.Context, msg IMessage, number int) func() {
	prepareQueue := fmt.Sprintf("%s.prepare", msg.GetChannel())

	if number <= 0 {
		number = 1
	}

	// the quit channel is used for waiting goroutines exit
	quit := make(chan struct{}, number)
	// the cancel slice store cancel functions for child contexts used in goroutines
	cancelSlice := make([]context.CancelFunc, number)

	for i := 0; i < number; i++ {
		// create child context each per goroutine
		childCtx, cancel := context.WithCancel(ctx)
		// store each cancel function of child contexts into cancel slice
		cancelSlice[i] = cancel

		go func(ctx context.Context, number int) {
			doingQueue := fmt.Sprintf("%s.doing%d", msg.GetChannel(), number)
			for {
				select {
				case <-ctx.Done():
					fmt.Println("context has been canceled")
					quit <- struct{}{}
					return
				default:
				}

				reply, msg, err := q.rpoplpush(msg, prepareQueue, doingQueue, true)
				if err != nil {
					fmt.Println("failed to pop msg", err)
					continue
				}
				if msg == nil {
					continue
				}
				if err := msg.Resolve(); err == nil {
					//time.Sleep(time.Second)
					q.lrem(doingQueue, reply)
				}
				q.ack(msg, doingQueue, prepareQueue)
			}
		}(childCtx, i) // the each variable i must be passed to goroutines immediately
	}

	cancelFunc := func() {
		for i := 0; i < number; i++ {
			cancel := cancelSlice[i]
			cancel()
			<-quit
		}
	}

	fmt.Printf("%d receivers have been initialized\n", number)
	return cancelFunc
}

// Delivery will send msg implementing IMessage into the queue to be consumed
func (q *Queue) Delivery(msg IMessage) error {
	conn := q.pool.Get()
	defer conn.Close()
	prepareQueue := fmt.Sprintf("%s.prepare", msg.GetChannel())
	if msgJson, err := msg.Marshal(); err != nil {
		return err
	} else {
		_, err := conn.Do("LPUSH", prepareQueue, msgJson)
		fmt.Println("produced", string(msgJson))

		return err
	}
}
