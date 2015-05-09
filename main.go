package main

import (
	"fmt"
	"time"
)

type ActorSystem struct {
	actors map[string]*ActorWrapper
}

type ActorWrapper struct {
	isRunning bool
	queue     chan interface{}
	actor     Actor
}

func (r *ActorWrapper) Send(message interface{}) {
	// TODO: handle Stop message
	r.queue <- message
}

func (r *ActorWrapper) Start() {
	fmt.Println("Starting actor")
	r.isRunning = true

	go func() {
		for r.isRunning {
			fmt.Println("Actor waiting for a message!")

			message := <-r.queue
			r.actor.receive(message)
		}
		fmt.Println("Actor exiting")
	}()
}

func (r *ActorWrapper) Stop() {
	r.isRunning = false
}

func (s *ActorSystem) AddActor(name string, actor Actor) *ActorWrapper {
	actorWrapper := new(ActorWrapper)
	actorWrapper.actor = actor
	actorWrapper.queue = make(chan interface{})

	s.actors[name] = actorWrapper

	return actorWrapper
}

func (s *ActorSystem) GetActor(name string) *ActorWrapper {
	return s.actors[name]
}

type Actor interface {
	receive(message interface{})
}

type CalculatorActor struct{}

type NumbersToSum struct {
	numbers []int
}

func (c *CalculatorActor) receive(message interface{}) {
	fmt.Println("CalculatorActor handling message")

	switch m := message.(type) {
	case NumbersToSum:
		{
			fmt.Println("We got numbers to sum")
			sum := 0
			for _, value := range m.numbers {
				sum = sum + value
			}
			fmt.Printf("\tSum: %d\n", sum)
		}
	}
}

type PrintlnActor struct{}

func (e *PrintlnActor) receive(message interface{}) {
	fmt.Println("PrintlnActor handling message")

	switch m := message.(type) {
	case string:
		{
			fmt.Println(m)
		}
	}
}

type PostmanActor struct {
	actorSystem *ActorSystem
}

type MessageToSend struct {
	recipient string
	contents  interface{}
}

func (p *PostmanActor) receive(message interface{}) {
	switch m := message.(type) {
	case MessageToSend:
		{
			p.SendMessage(m)
		}
	}
}

func (m *PostmanActor) SendMessage(message MessageToSend) {
	fmt.Printf("Trying to send a message to %s\n", message.recipient)

	recipient := m.actorSystem.GetActor(message.recipient)
	if recipient != nil {
		recipient.Send(message.contents)
	} else {
		fmt.Println("Recipient does not exist")
	}
}

func CreateActorSystem() *ActorSystem {
	actorSystem := new(ActorSystem)
	actorSystem.actors = make(map[string]*ActorWrapper)
	return actorSystem
}

func (s *ActorSystem) Length() int {
	return len(s.actors)
}

func main() {
	actorSystem := CreateActorSystem()
	fmt.Printf("Length: %d\n", actorSystem.Length())

	calculatorActor := new(CalculatorActor)
	calculatorActorWrapper := actorSystem.AddActor("calculator1", calculatorActor)
	fmt.Printf("Length: %d\n", actorSystem.Length())

	printlnActor := new(PrintlnActor)
	printlnActorWrapper := actorSystem.AddActor("println1", printlnActor)
	fmt.Printf("Length: %d\n", actorSystem.Length())

	postmanActor := new(PostmanActor)
	postmanActor.actorSystem = actorSystem
	postmanActorWrapper := actorSystem.AddActor("postman1", postmanActor)

	calculatorActorWrapper.Start()
	printlnActorWrapper.Start()
	postmanActorWrapper.Start()

	printlnActorWrapper.Send("hello!")
	calculatorActorWrapper.Send(NumbersToSum{[]int{1, 2, 3, 4}})
	calculatorActorWrapper.Send(NumbersToSum{[]int{4, 5, 6, 7}})
	postmanActorWrapper.Send(MessageToSend{"calculator1", NumbersToSum{[]int{4, 5, 6, 7, 1}}})
	printlnActorWrapper.Send("hello again!")

	calculatorActorWrapper.Stop()
	printlnActorWrapper.Stop()
	postmanActorWrapper.Stop()

	time.Sleep(500 * time.Millisecond)

	fmt.Println("Leaving")
}
