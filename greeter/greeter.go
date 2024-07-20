package greeter

import (
	"fmt"

	"github.com/google/wire"
)

type Message string

// provider function
func NewMessage() Message {
	return Message("Hi there!")
}

type Greeter struct {
	Message Message
}

func (g Greeter) Greet() Message {
	return g.Message
}

// provider function
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

type Event struct {
	Greeter Greeter
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

var MainBindingSet = wire.NewSet(NewMessage, NewGreeter, NewEvent)
