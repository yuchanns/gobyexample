package newbee_traps

import "fmt"

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
