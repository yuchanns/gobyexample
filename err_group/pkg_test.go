package err_group

import (
	"errors"
	"fmt"
	"runtime"
	"testing"
)

var goprocs = runtime.GOMAXPROCS(0)

var (
	start = 300
	end   = 600
	step  = 10
)

func BenchmarkChanErrGroup(b *testing.B) {
	for i := start; i < end; i += step {
		errGroup := NewChanErrGroup()
		b.Run(fmt.Sprintf("goroutine-%d", i*goprocs), func(b *testing.B) {
			b.SetParallelism(i)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					errGroup.Add(errors.New("a"))
				}
			})
		})
	}
}

func BenchmarkLockErrGroup(b *testing.B) {
	for i := start; i < end; i += step {
		errGroup := NewLockErrGroup()
		b.Run(fmt.Sprintf("goroutine-%d", i*goprocs), func(b *testing.B) {
			b.SetParallelism(i)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					errGroup.Add(errors.New("a"))
				}
			})
		})
	}
}
