package chain_pattern

import (
	"context"
	"fmt"
)

type NotifyConsumer struct {
	*AConsumer
}

func NewNotifyCnosumer() *AConsumer {
	return NewAConsumerWithConsumer(TagNotify, &NotifyConsumer{})
}

func (c *NotifyConsumer) Consume(_ context.Context, msg *Message) error {
	fmt.Printf("notify consumer consumes message: %v\n", msg.Data)
	return nil
}
