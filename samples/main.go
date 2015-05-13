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
	types := []interface{}{NumbersToSum{}, NumbersToMultiply{}}

	actorSystem.AddActor("calculator", types, &CalculatorActor{1}).Start()
	actorSystem.AddActor("calculator", types, &CalculatorActor{2}).Start()

	actorSystem.SendMessage(NumbersToSum{[]int{1, 2, 3, 4}})
	actorSystem.SendMessage(NumbersToMultiply{[]int{1, 2, 3, 4}})
	actorSystem.SendMessage(NumbersToDivide{[]int{1, 2, 3, 4}})

	time.Sleep(1 * time.Second)

	actorSystem.Stop()

	fmt.Println("Leaving")
}
