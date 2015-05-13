package chewbakka

import (
	"fmt"
)

type DeadMessagesActor struct{}

func (d *DeadMessagesActor) Receive(message interface{}) {
	fmt.Println("Received a message but it was lost")
	// TODO: handle lost message
}
