package main

import (
	"fmt"
	"testing"
)

func BenchmarkFib(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	_ = fib(999999)
}

func BenchmarkMemoryLeak(b *testing.B) {
	memoryLeak()
	_ = fmt.Sprintln(leakSlice)
}
