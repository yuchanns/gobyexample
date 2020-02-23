package base

import (
	"fmt"
	"math/rand"
)

type Vehicle interface {
	Load(int)
	Run()
}

type Car struct {
	num int
}

func (c *Car) Load(num int) {
	c.num = num
	fmt.Printf("the car loaded %d people on the road side.\n", num)
}

func (Car) Run() {
	fmt.Println("the car is running on the road.")
}

type Plan struct {
	num int
}

func (p *Plan) Load(num int) {
	p.num = num
	fmt.Printf("the plan loaded %d people by the ladder\n", num)
}

func (Plan) Run() {
	fmt.Println("the plan is flying in the sky.")
}

func Station(v Vehicle) {
	v.Load(rand.Intn(10))
	v.Run()
}
