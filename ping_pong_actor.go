package main

import (
	"fmt"
)

type PingPongActor struct {
	actorSystem     *ActorSystem
	otherPlayerName string
}

type PingPongBall struct {
	count int
	max   int
}

func (a *PingPongActor) receive(message interface{}) {
	switch m := message.(type) {
	case PingPongBall:
		if m.count <= m.max {
			if m.count%2 == 0 {
				fmt.Printf("Ping! %d\n", m.count)
			} else {
				fmt.Printf("Pong! %d\n", m.count)
			}

			otherPlayer := a.actorSystem.GetActor(a.otherPlayerName)
			m.count = m.count + 1
			otherPlayer.Send(m)
		} else {
			fmt.Println("Match ended!")
		}
	}
}
