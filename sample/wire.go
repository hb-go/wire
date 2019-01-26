// +build wireinject

package sample

import (
	"errors"
	"fmt"
	"github.com/google/wire"
	"time"
)

func NewEventNumber() int {
	return 1
}

func InitializeEvent(phrase string) (Event, error) {
	// wire.Build(NewEvent, NewMessage)
	// wire.Build(NewEvent, NewGreeter, NewMessage, NewEventNumber)
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}

type Message string

type Greeter struct {
	Message Message
	Grumpy  bool
}

type Event struct {
	Greeter Greeter
}

func NewMessage(phrase string) (Message, error) {
	if len(phrase) <= 0 {
		return "", errors.New("could not create message: phrase length must > 0")
	}
	return Message(phrase), nil
}

func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
