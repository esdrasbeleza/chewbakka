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
	types := []interface{}{NumbersToMultiply{}, NumbersToSum{}}
	calculatorActorWrapper := actorSystem.AddActor("calculator1", types, calculatorActor)
	calculatorActorWrapper.Start()

	actorSystem.SendMessage(NumbersToSum{[]int{1, 2, 3, 4}})
	actorSystem.SendMessage("This message won't be handled")
	actorSystem.SendMessage(NumbersToMultiply{[]int{1, 2, 3, 4}})

	time.Sleep(4 * time.Second)

	calculatorActorWrapper.Stop()

	fmt.Println("Leaving")
}
