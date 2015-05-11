package main

import (
	"fmt"
	"time"

	"github.com/esdrasbeleza/chewbakka"
)

func main() {
	// Create actor system
	actorSystem := chewbakka.CreateActorSystem()

	// Running Calculator Actor
	calculatorActor := new(CalculatorActor)
	calculatorActorWrapper := actorSystem.AddActor("calculator1", calculatorActor)
	calculatorActorWrapper.Start()
	calculatorActorWrapper.Send(NumbersToSum{[]int{1, 2, 3, 4}})
	calculatorActorWrapper.Send(NumbersToMultiply{[]int{1, 2, 3, 4}})

	// Running Postman Actor
	postmanActor := new(PostmanActor)
	postmanActor.actorSystem = actorSystem
	postmanActorWrapper := actorSystem.AddActor("postman1", postmanActor)
	postmanActorWrapper.Start()
	postmanActorWrapper.Send(MessageToSend{"calculator1", NumbersToSum{[]int{4, 5, 6, 7, 1}}})

	// Running Ping Pong actors
	pingActor := PingPongActor{actorSystem: actorSystem, otherPlayerName: "pong"}
	pongActor := PingPongActor{actorSystem: actorSystem, otherPlayerName: "ping"}

	pingWrapper := actorSystem.AddActor("ping", &pingActor)
	pongWrapper := actorSystem.AddActor("pong", &pongActor)

	pingWrapper.Start()
	pongWrapper.Start()

	pingWrapper.Send(PingPongBall{0, 10})

	// Running Println Actor
	printlnActor := new(PrintlnActor)
	printlnActorWrapper := actorSystem.AddActor("println1", printlnActor)
	fmt.Printf("Length: %d\n", actorSystem.Length())

	printlnActorWrapper.Start()
	printlnActorWrapper.Send("hello, stopped actor!")
	printlnActorWrapper.Send("hello, stopped actor again!")

	time.Sleep(2 * time.Second)
	printlnActorWrapper.Stop()

	pingWrapper.Stop()
	pongWrapper.Stop()
	calculatorActorWrapper.Stop()
	postmanActorWrapper.Stop()

	time.Sleep(2 * time.Second)

	fmt.Println("Leaving")
}
