package error_trace

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"runtime"
	"strings"
)

func ErrHandle(f func(ctx context.Context) error) func(ctx context.Context) {
	return func(ctx context.Context) {
		defer func() {
			msg := recover()
			if msg != nil {
				switch err := msg.(type) {
				case runtime.Error:
					stackErr := errors.Wrapf(err, "panic runtime error: %v", err)
					PrintStack(stackErr, 4)
				default:
					stackErr := errors.New(fmt.Sprintf("panic error: %v", err))
					PrintStack(stackErr, 4)
				}
			}
		}()
		err := f(ctx)
		if err != nil {
			PrintStack(err, 0)
		}
	}
}

func PrintStack(err error, skip int) {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}
	if err, ok := err.(stackTracer); ok {
		frames := err.StackTrace()
		l := len(frames)
		traces := make([]string, 0)
		for i := skip; i <= skip+4; i++ {
			if i >= l {
				break
			}

			f := frames[i]
			traces = append(traces, strings.ReplaceAll(strings.TrimSpace(fmt.Sprintf("%+s:%d", f, f)), "\n\t", " "))
		}
		fmt.Printf("%s\n", err)
		fmt.Printf("error stack: %v\n", traces)
	}
}
