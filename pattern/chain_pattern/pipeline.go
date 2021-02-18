package chain_pattern

import (
	"context"
	"errors"
)

// PipeLine holds a chain of real consumers and tells them to consume msg.
type PipeLine struct {
	firstConsumer *AConsumer
}

func NewPipline(consumers ...*AConsumer) *PipeLine {
	p := &PipeLine{}
	l := len(consumers)
	// build the chain of responsibility
	if l > 0 {
		consumer := consumers[0]
		p.firstConsumer = consumer
		for i := 1; i < l; i++ {
			consumers[i-1].SetNext(consumers[i])
		}
	}
	return p
}

func (p *PipeLine) Consume(ctx context.Context, msg *Message) error {
	if p.firstConsumer == nil {
		return errors.New("first consumer not found")
	}
	return p.firstConsumer.Consume(ctx, msg)
}
