package main

import "fmt"

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
