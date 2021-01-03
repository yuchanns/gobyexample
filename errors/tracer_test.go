package errors

import (
	"context"
	"github.com/pkg/errors"
	"sync"
	"testing"
)

func TestBuildStack(t *testing.T) {
	logger := BuildLogger()
	err1 := errors.New("first error")
	err2 := errors.Wrap(err1, "second error")
	stacks := BuildStack(err2, 0)
	logger.Err(err2).Fields(map[string]interface{}{"stack": stacks}).Msg("errors happen")
}

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
	wg := &sync.WaitGroup{}
	wg.Add(len(tables))
	for i := range tables {
		go func(i int) {
			defer wg.Done()
			PanicHandle(tables[i])(context.TODO())
		}(i)
	}
	wg.Wait()
}
