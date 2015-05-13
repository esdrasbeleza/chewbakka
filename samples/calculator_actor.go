package main

import (
	"fmt"
)

type CalculatorActor struct {
	Id int
}

type NumbersToSum struct {
	numbers []int
}

type NumbersToMultiply struct {
	numbers []int
}

type NumbersToDivide struct {
	numbers []int
}

func (c *CalculatorActor) Receive(message interface{}) {
	fmt.Printf("CalculatorActor (id:%v) handling message\n", c.Id)

	switch m := message.(type) {
	case NumbersToSum:
		{
			fmt.Printf("CalculatorActor (id:%v) We got numbers to sum\n", c.Id)
			sum := 0
			for _, value := range m.numbers {
				sum = sum + value
			}
			fmt.Printf("\tResult: %d\n", sum)
		}
	case NumbersToMultiply:
		{
			fmt.Printf("CalculatorActor (id:%v) We got numbers to multiply\n", c.Id)
			sum := 1
			for _, value := range m.numbers {
				sum = sum * value
			}
			fmt.Printf("\tResult: %d\n", sum)
		}
	}
}
