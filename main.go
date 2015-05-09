package main

import (
	"fmt"
)

type ActorSystem struct {
	actors map[string]*ActorWrapper
}

type ActorWrapper struct {
	isRunning bool
	queue     interface{} // TODO: implement queue here
	actor     Actor
}

func (r *ActorWrapper) Send(message interface{}) {
	// TODO: enqueue message
	r.actor.receive(message)
}

func (r *ActorWrapper) Start() {
	// TODO: create message queue
	// TODO: set isRunning to true?
}

func (s *ActorSystem) AddActor(name string, actor Actor) *ActorWrapper {
	actorWrapper := new(ActorWrapper)
	actorWrapper.actor = actor

	// TODO: what to do if name already exists?
	s.actors[name] = actorWrapper

	return actorWrapper
}

func (s *ActorSystem) GetActor(name string) *ActorWrapper {
	// TODO: what if the actor does not exist?
	return s.actors["name"]
}

type Actor interface {
	receive(message interface{})
}

type CalculatorActor struct {
}

type NumbersToSum struct {
	numbers []int
}

func (c *CalculatorActor) receive(message interface{}) {
	switch m := message.(type) {
	case NumbersToSum:
		{
			sum := 0
			for _, value := range m.numbers {
				sum = sum + value
			}
			fmt.Printf("Sum: %d", sum)
		}
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

	calculator := new(CalculatorActor)
	actorRef := actorSystem.AddActor("calculator1", calculator)
	fmt.Printf("Length: %d\n", actorSystem.Length())

	actorRef.Send("hello")
}
