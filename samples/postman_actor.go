package main

import (
	"fmt"

	"github.com/esdrasbeleza/chewbakka"
)

type PostmanActor struct {
	actorSystem *chewbakka.ActorSystem
}

type MessageToSend struct {
	recipient string
	contents  interface{}
}

func (p *PostmanActor) Receive(message interface{}) {
	switch m := message.(type) {
	case MessageToSend:
		{
			p.SendMessage(m)
		}
	}
}

func (m *PostmanActor) SendMessage(message MessageToSend) {
	fmt.Printf("Trying to send a message to %s\n", message.recipient)

	recipient := m.actorSystem.GetActor(message.recipient)
	if recipient != nil {
		recipient.Send(message.contents)
	} else {
		fmt.Println("Recipient does not exist")
	}
}
