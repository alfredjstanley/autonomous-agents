# autonomous-agents

A lightweight actor-based system inspired by emergent intelligence and distributed messaging.

## 💡 Concept

Each **Agent** is:

- A goroutine with a message inbox
- Able to receive, store, and forward messages
- Controlled by a central **Controller**

## 📦 Structure

- `agent/`: core actor logic
- `controller/`: system orchestrator
- `message/`: inter-agent message protocol

## 🚀 Getting Started

```bash
git clone https://github.com/alfredjstanley/autonomous-agents
cd autonomous-agents; go run cmd/main.go
```
