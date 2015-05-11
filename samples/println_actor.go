package main

import (
	"fmt"
)

type PrintlnActor struct{}

func (e *PrintlnActor) Receive(message interface{}) {
	fmt.Println("PrintlnActor handling message")

	switch m := message.(type) {
	case string:
		{
			fmt.Println(m)
		}
	}
}
