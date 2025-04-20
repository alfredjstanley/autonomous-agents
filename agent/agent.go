package agent

import (
	"fmt"

	"github.com/alfredjstanley/autonomous-agents/message"
)

type Agent struct {
	ID      string
	Inbox   chan message.Message
	Memory  []string
	Running bool
}

func NewAgent(id string) *Agent {
	a := &Agent{
		ID:      id,
		Inbox:   make(chan message.Message),
		Memory:  []string{},
		Running: true,
	}
	go a.Run()
	return a
}

func (a *Agent) Run() {
	fmt.Printf("[Agen %s] Online. \n", a.ID)
	for a.Running {
		select {
		case msg := <-a.Inbox:
			a.handleMessage(msg)
		default:
			fmt.Println("default case")
		}

	}
	fmt.Printf("[Agen %s] Offline. \n", a.ID)

}

func (a *Agent) handleMessage(msg message.Message) {
	switch msg.Type {
	case message.MsgPing:
		fmt.Printf("[Agent %s] Received PING from %s\n", a.ID, msg.From)
	case message.MsgStoreMemory:
		fmt.Printf("[Agent %s] Storing: %s\n", a.ID, msg.Payload)
		a.Memory = append(a.Memory, msg.Payload)
	case message.MsgForward:
		fmt.Printf("[Agent %s] FORWARDING message to %s\n", a.ID, msg.Payload)
	case message.MsgShutdown:
		fmt.Printf("[Agent %s] SHUTTING DOWN\n", a.ID)
		a.Running = false
	}
}
