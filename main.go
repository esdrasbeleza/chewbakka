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
