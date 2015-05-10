package main

import (
	"fmt"
	"time"
)

func main() {
	actorSystem := CreateActorSystem()
	fmt.Printf("Length: %d\n", actorSystem.Length())

	calculatorActor := new(CalculatorActor)
	calculatorActorWrapper := actorSystem.AddActor("calculator1", calculatorActor)
	fmt.Printf("Length: %d\n", actorSystem.Length())

	postmanActor := new(PostmanActor)
	postmanActor.actorSystem = actorSystem
	postmanActorWrapper := actorSystem.AddActor("postman1", postmanActor)

	calculatorActorWrapper.Start()
	postmanActorWrapper.Start()

	calculatorActorWrapper.Send(NumbersToSum{[]int{1, 2, 3, 4}})
	calculatorActorWrapper.Send(NumbersToSum{[]int{4, 5, 6, 7}})
	postmanActorWrapper.Send(MessageToSend{"calculator1", NumbersToSum{[]int{4, 5, 6, 7, 1}}})

	// Sample: ping pong actors
	pingActor := PingPongActor{actorSystem: actorSystem, otherPlayerName: "pong"}
	pongActor := PingPongActor{actorSystem: actorSystem, otherPlayerName: "ping"}

	pingWrapper := actorSystem.AddActor("ping", &pingActor)
	pongWrapper := actorSystem.AddActor("pong", &pongActor)

	pingWrapper.Start()
	pongWrapper.Start()

	pingWrapper.Send(PingPongBall{0, 10})

	// Sample: println actor
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
