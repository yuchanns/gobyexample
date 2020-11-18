package error_trace

import (
	"context"
	"github.com/pkg/errors"
	"testing"
)

func TestPanicHandle(t *testing.T) {
	tables := []func(ctx context.Context) error{
		func(ctx context.Context) error {
			var err error
			err.Error()
			return nil
		},
		func(ctx context.Context) error {
			return errors.New("a regular error")
		},
		func(ctx context.Context) error {
			panic("a panic error")
		},
	}
	for i := range tables {
		ErrHandle(tables[i])(context.TODO())
	}
}
