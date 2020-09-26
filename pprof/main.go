package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func fibIter(a, b, n int) int {
	if n == 0 {
		return b
	}

	return fibIter(a+b, a, n-1)
}

func fib(n int) int {
	return fibIter(1, 0, n)
}

func main() {
	f, err := os.Create("cpu.out")
	if err != nil {
		log.Fatal("failed to create cpu profile:", err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("failed to start cpu profile:", err)
	}
	defer pprof.StopCPUProfile()
	result := fib(999999)
	fmt.Println(result)
}
