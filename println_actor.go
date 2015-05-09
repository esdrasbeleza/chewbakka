package main

import "fmt"

type PrintlnActor struct{}

func (e *PrintlnActor) receive(message interface{}) {
	fmt.Println("PrintlnActor handling message")

	switch m := message.(type) {
	case string:
		{
			fmt.Println(m)
		}
	}
}
