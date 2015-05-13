package chewbakka

import (
	"fmt"
	"reflect"
)

type Actor interface {
	Receive(message interface{})
}

type DeadLetters struct{}

func (c *DeadLetters) Receive(message interface{}) {
	fmt.Println("** Received a Dead Letter **\n", reflect.TypeOf(message), message)
}

type ActorRouter interface {
	Select([]*ActorWrapper) *ActorWrapper
}

type RoundRobin struct {
	position int
}

func (x *RoundRobin) Select(v []*ActorWrapper) *ActorWrapper {
	x.position += 1
	return v[x.position%len(v)]
}

type ActorSystem struct {
	actors      map[string][]*ActorWrapper
	recieves    map[string][]reflect.Type
	router      ActorRouter
	deadLetters *ActorWrapper
}

type ActorWrapper struct {
	isRunning   bool
	queue       chan interface{}
	actor       Actor
	actorSystem *ActorSystem
	actorName   string
	messageType []reflect.Type
}

func (r *ActorWrapper) Send(message interface{}) {
	// Avoid blocking the main thread
	go func() {
		if r.isRunning {
			r.queue <- message
		}
	}()
}

func (r *ActorWrapper) Start() {
	fmt.Println("Starting actor")
	r.isRunning = true

	go func() {
		for r.isRunning {
			fmt.Println("Actor waiting for a message!")

			message, ok := <-r.queue // Will block until we have a message

			// Since the value may have changed, we check it again
			if r.isRunning && ok {
				r.actor.Receive(message)
			}
		}

		fmt.Println("Actor left the stage")
	}()
}

func (r *ActorWrapper) Stop() {
	fmt.Println("Stopping actor")
	r.isRunning = false
	close(r.queue)
	r.actorSystem.RemoveActor(r.actorName)

	fmt.Printf("Actor system now has %d actors\n", r.actorSystem.Length())
}

func CreateActorSystem() *ActorSystem {
	actorSystem := new(ActorSystem)
	actorSystem.actors = make(map[string][]*ActorWrapper)
	actorSystem.recieves = make(map[string][]reflect.Type)
	actorSystem.router = new(RoundRobin)

	deadLetters := new(ActorWrapper)
	deadLetters.actor = new(DeadLetters)
	deadLetters.queue = make(chan interface{})
	deadLetters.actorSystem = actorSystem
	deadLetters.actorName = "DeadLetters"
	deadLetters.messageType = []reflect.Type{}

	actorSystem.deadLetters = deadLetters

	deadLetters.Start()

	return actorSystem
}

func (s *ActorSystem) AddActor(name string, messageTypes []interface{}, actor Actor) *ActorWrapper {
	types := make([]reflect.Type, len(messageTypes))
	for i, v := range messageTypes {
		types[i] = reflect.TypeOf(v)
	}

	matches := func(x, y []reflect.Type) bool {
		if len(x) != len(y) {
			return false
		}
		for k, v := range x {
			if v != y[k] {
				return false
			}
		}
		return true
	}

	x, ok := s.recieves[name]
	if !ok || matches(types, x) {

		actorWrapper := new(ActorWrapper)
		actorWrapper.actor = actor
		actorWrapper.queue = make(chan interface{})
		actorWrapper.actorSystem = s
		actorWrapper.actorName = name
		actorWrapper.messageType = types

		if x, ok := s.actors[name]; ok {
			s.actors[name] = append(x, actorWrapper)
		} else {
			s.actors[name] = []*ActorWrapper{actorWrapper}
		}
		s.recieves[name] = types

		return actorWrapper
	}
	return nil
}

func (s *ActorSystem) RemoveActor(name string) {
	delete(s.actors, name)
}

func (s *ActorSystem) SendMessage(message interface{}) {
	messageType := reflect.TypeOf(message)

	contains := func(x []reflect.Type, y reflect.Type) bool {
		for _, v := range x {
			if v == y {
				return true
			}
		}
		return false
	}

	found := false
	for k, v := range s.recieves {
		if contains(v, messageType) {
			if x, ok := s.actors[k]; ok {
				found = true

				actor := s.router.Select(x)
				actor.Send(message)
			}
		}
	}

	if !found {
		s.deadLetters.Send(message)
	}
}

func (s *ActorSystem) GetActors(name string) []*ActorWrapper {
	return s.actors[name]
}

func (s *ActorSystem) Stop() {
	for _, x := range s.actors {
		for _, y := range x {
			y.Stop()
		}
	}
}

func (s *ActorSystem) Length() int {
	sum := 0
	for _, x := range s.actors {
		sum += len(x)
	}
	return sum
}
