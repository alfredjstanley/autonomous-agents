package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/alfredjstanley/autonomous-agents/controller"
	"github.com/alfredjstanley/autonomous-agents/message"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	c := controller.NewController()

	// Create agents
	for i := 0; i < 3; i++ {
		id := fmt.Sprintf("agent-%d", i+1)
		c.AddAgent(id)
	}

	// Broadcast ping
	c.Broadcast(message.MsgPing, "")

	// Send memory messages
	c.SendMessage("controller", "agent-1", message.MsgStoreMemory, "Hello from ECHO")
	c.SendMessage("controller", "agent-2", message.MsgStoreMemory, "Observe and adapt")

	time.Sleep(1 * time.Second)

	// Shutdown all
	c.ShutdownAll()
	time.Sleep(500 * time.Millisecond)

}
