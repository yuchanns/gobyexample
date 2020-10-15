package wire

import "fmt"

type Message string

type Greeter struct {
	message Message
}

type Event struct {
	greeter *Greeter
}

func NewMessage() Message {
	return "hello wire!"
}

func NewGreeter(m Message) *Greeter {
	return &Greeter{message: m}
}

func NewEvent(g *Greeter) *Event {
	return &Event{greeter: g}
}

func (g *Greeter) Greet() Message {
	return g.message
}

func (e *Event) Start() {
	msg := e.greeter.Greet()
	fmt.Println(msg)
}
