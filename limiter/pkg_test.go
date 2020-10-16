package limiter

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNewTokenBucket(t *testing.T) {
	tb := NewTokenBucket(2, 1)
	wg := &sync.WaitGroup{}
	counts := 10
	wg.Add(counts)
	for i := 0; i < counts; i++ {
		go func(i int) {
			for {
				if tb.Allow() {
					fmt.Printf("%d pass, time is %s\n", i, time.Now().Format("2006-01-02 15:04:05"))
					wg.Done()
					return
				}
			}
		}(i)
	}
	wg.Wait()
}
