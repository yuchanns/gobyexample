package chain_pattern

import (
	"context"
	"fmt"
)

type OrderConsumer struct {
	*AConsumer
}

func NewOrderConsumer() *AConsumer {
	return NewAConsumerWithConsumer(TagOrder, &OrderConsumer{})
}

func (c *OrderConsumer) Consume(_ context.Context, msg *Message) error {
	fmt.Printf("order consumer consumes message: %v\n", msg.Data)
	return nil
}
