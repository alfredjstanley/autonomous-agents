package controller

import (
	"fmt"

	"github.com/alfredjstanley/autonomous-agents/agent"
	"github.com/alfredjstanley/autonomous-agents/message"
)

type Controller struct {
	Agents map[string]*agent.Agent
}

func NewController() *Controller {
	return &Controller{
		Agents: make(map[string]*agent.Agent),
	}
}

func (c *Controller) AddAgent(id string) {
	a := agent.NewAgent(id)
	c.Agents[id] = a
}

func (c *Controller) Broadcast(msgType message.MessageType, payload string) {
	for id := range c.Agents {
		c.SendMessage("controller", id, msgType, payload)
	}
}

func (c *Controller) SendMessage(from, to string, msgType message.MessageType, payload string) {
	target, exists := c.Agents[to]
	if !exists {
		fmt.Printf("[Controller] Agent %s is not found\n", to)
	}

	target.Inbox <- message.Message{
		From:    from,
		To:      to,
		Type:    msgType,
		Payload: payload,
	}
}

func (c *Controller) ShutdownAll() {
	for id := range c.Agents {
		c.SendMessage("controller", id, message.MsgShutdown, "")
	}
}
