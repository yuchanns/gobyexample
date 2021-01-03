package errors

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"os"
	"runtime"
	"strings"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func BuildLogger() *zerolog.Logger {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &logger
}

func BuildStack(err error, skip int) []string {
	traces := make([]string, 0)
	if err, ok := err.(stackTracer); ok {
		frames := err.StackTrace()
		l := len(frames)
		for i := skip; i <= skip+4; i++ {
			if i >= l {
				break
			}
			f := frames[i]
			traces = append(traces, strings.ReplaceAll(
				strings.TrimSpace(fmt.Sprintf("%+s:%d", f, f)),
				"\n\t",
				" ",
			))
		}
	}
	return traces
}

func PanicHandle(f func(ctx context.Context) error) func(ctx context.Context) {
	logger := BuildLogger()
	return func(ctx context.Context) {
		defer func() {
			msg := recover()
			if msg != nil {
				var stacks []string
				var stackErr error
				switch err := msg.(type) {
				case runtime.Error:
					stackErr = errors.Wrap(err, "panic runtime error: ")
					stacks = BuildStack(stackErr, 4)
				default:
					stackErr = errors.New(fmt.Sprintf("panic error: %v", err))
					stacks = BuildStack(stackErr, 2)
				}
				logger.Err(stackErr).Fields(map[string]interface{}{"stacks": stacks}).Msg(stackErr.Error())
			}
		}()
		err := f(ctx)
		if err != nil {
			stacks := BuildStack(err, 0)
			logger.Err(err).Fields(map[string]interface{}{"stacks": stacks}).Msg(err.Error())
		}
	}
}
