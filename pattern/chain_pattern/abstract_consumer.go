package chain_pattern

import "context"

// IConsumer is the base interface of all real consumers.
type IConsumer interface {
	// Consume requires an implementation that can consume the given msg.
	// Normally, we use a specific struct represent the argument msg instead of Message
	// but now there is no need to be so detailed.
	Consume(ctx context.Context, msg *Message) error
}

// AConsumer is the base abstract class of real consumers.
type AConsumer struct {
	// With next we can build the chain of real consumers.
	next *AConsumer
	// Tag decides a real consumer's responsibility.
	Tag Tag
	// IConsumer is embedded in order to use the methods implemented by real consumers.
	IConsumer
}

// NewAConsumerWithConsumer returns a real consumer implemented based on AConsumer
func NewAConsumerWithConsumer(tag Tag, c IConsumer) *AConsumer {
	return &AConsumer{Tag: tag, IConsumer: c}
}

// SetNext sets a consumer after current consumer on the chain.
func (c *AConsumer) SetNext(consumer *AConsumer) {
	c.next = consumer
}

// Consume tells real consumers to consume the msg
func (c *AConsumer) Consume(ctx context.Context, msg *Message) error {
	if c.Tag == msg.Tag {
		return c.IConsumer.Consume(ctx, msg)
	}
	return c.next.Consume(ctx, msg)
}
