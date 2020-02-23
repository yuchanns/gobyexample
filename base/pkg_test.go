package base

import (
	"fmt"
	"testing"
)

func TestUserAges(t *testing.T) {
	ua := UserAges{ages: make(map[string]int)}
	for i := 0; i < 10000; i++ {
		go func(i int) {
			ua.Add("yuchanns", i)
		}(i)
	}
	for i := 0; i < 10000; i++ {
		go func() {
			println(ua.Get("yuchanns"))
		}()
	}
}

func TestNilMap(t *testing.T) {
	m := make(map[string]int)
	x := m
	m["hello"] = 42
	fmt.Println(x["hello"])
}

func TestStation(t *testing.T) {
	c := &Car{}
	p := &Plan{}
	Station(c)
	Station(p)
}
