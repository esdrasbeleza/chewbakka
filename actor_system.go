package main

import "fmt"

type ActorSystem struct {
	actors map[string]*ActorWrapper
}

type ActorWrapper struct {
	isRunning bool
	queue     chan interface{}
	actor     Actor
}

func (r *ActorWrapper) Send(message interface{}) {
	r.queue <- message
}

func (r *ActorWrapper) Start() {
	fmt.Println("Starting actor")
	r.isRunning = true

	go func() {
		for r.isRunning {
			fmt.Println("Actor waiting for a message!")

			message := <-r.queue // Will block until we have a message, fuck yeah
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
