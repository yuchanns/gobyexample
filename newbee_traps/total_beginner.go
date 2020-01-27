package newbee_traps

import (
	"fmt"
	"time"
	"unicode/utf8"
)

func NilInitVariableWithExplicitType() interface{} {
	var x interface{} = nil
	return x
}

func NilInitSlicesAndMaps() (map[string]int, []int) {
	m := make(map[string]int)
	m["one"] = 1

	var s []int
	s = append(s, 1)

	return m, s
}

func InitStrings() string {
	var s string

	if s == "" {
		s = "default"
	}

	return s
}

func RangeSlices() {
	x := []string{"a", "b", "c"}

	for _, v := range x {
		fmt.Println(v)
	}
}

func MultiDimension() [][]int {
	x := 2
	y := 4

	table := make([][]int, x)
	for i := range table {
		table[i] = make([]int, y)
	}

	fmt.Println(table)

	return table
}

func ImmutableStrings() (string, string) {
	x := "test"
	xbytes := []byte(x)
	xbytes[0] = 'T'
	y := "sj"
	yrunes := []rune(y)
	yrunes[0] = '世'
	yrunes[1] = '界'

	fmt.Println(x[0])

	for _, v := range yrunes {
		fmt.Printf("%#x\n", v)
	}

	return string(xbytes), string(yrunes)
}

func ValidateStringAndLength(data string) (bool, int, int) {
	result := utf8.ValidString(data)
	length := len(data)
	cLength := utf8.RuneCountInString(data)

	return result, length, cLength
}

func NilChannel() {
	inCh := make(chan int)
	outCh := make(chan int)

	go func() {
		var in <-chan int = inCh
		var out chan<- int
		var val int

		for {
			select {
			case out <- val:
				println("--------")
				out = nil
				in = inCh
			case val = <-in:
				println("++++++++++")
				out = outCh
				in = nil
			}
		}
	}()

	go func() {
		for r := range outCh {
			fmt.Println("Result: ", r)
		}
	}()

	time.Sleep(0)
	inCh <- 1
	inCh <- 2
	time.Sleep(3 * time.Second)
}
