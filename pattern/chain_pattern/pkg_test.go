package chain_pattern

import (
	"context"
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
			_ = pipeline.Consume(context.TODO(), message)
		}(message)
	}
	wg.Wait()
}
