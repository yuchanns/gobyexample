package chain_pattern

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

func TestNewPipline(t *testing.T) {
	orderConsumer := NewOrderConsumer()
	notifyConsumer := NewNotifyCnosumer()

	pipeline := NewPipline(orderConsumer, notifyConsumer)

	table := []*Message{
		{
			Tag:  TagOrder,
			Data: "order 1",
		},
		{
			Tag:  TagNotify,
			Data: "notify 1",
		},
		{
			Tag:  TagOrder,
			Data: "order 2",
		},
		{
			Tag:  TagNotify,
			Data: "notify 2",
		},
	}
	wg := &sync.WaitGroup{}
	wg.Add(len(table))
	for _, message := range table {
		go func(message *Message) {
			defer wg.Done()
			if err := pipeline.Consume(context.TODO(), message); err != nil {
				t.Error(err)
			}
		}(message)
	}
	wg.Wait()
}

func TestConsumerNotFound(t *testing.T) {
	orderConsumer := NewOrderConsumer()
	pipeline := NewPipline(orderConsumer)
	table := []*Message{
		{
			Tag:  TagNotify,
			Data: "notify 1",
		},
		{
			Tag:  TagNotify,
			Data: "notify 2",
		},
	}
	wg := &sync.WaitGroup{}
	wg.Add(len(table))
	for _, message := range table {
		go func(message *Message) {
			defer wg.Done()
			if err := pipeline.Consume(context.TODO(), message); err != nil {
				if err.Error() != fmt.Errorf("no consumer found for tag %s", message.Tag).Error() {
					t.Errorf("unexpected error: %s", err)
				}
			}
		}(message)
	}
	wg.Wait()
}
