package chewbakka

import (
	"fmt"
)

type Actor interface {
	Receive(message interface{})
}

type ActorSystem struct {
	actors map[string]*ActorWrapper
}

type ActorWrapper struct {
	isRunning   bool
	queue       chan interface{}
	actor       Actor
	actorSystem *ActorSystem
	actorName   string
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
	actorSystem.actors = make(map[string]*ActorWrapper)
	return actorSystem
}

func (s *ActorSystem) AddActor(name string, actor Actor) *ActorWrapper {
	actorWrapper := new(ActorWrapper)
	actorWrapper.actor = actor
	actorWrapper.queue = make(chan interface{})
	actorWrapper.actorSystem = s
	actorWrapper.actorName = name

	s.actors[name] = actorWrapper

	return actorWrapper
}

func (s *ActorSystem) RemoveActor(name string) {
	delete(s.actors, name)
}

func (s *ActorSystem) GetActor(name string) *ActorWrapper {
	return s.actors[name]
}

func (s *ActorSystem) Length() int {
	return len(s.actors)
}
