package message

type MessageType int

const (
	MsgPing MessageType = iota
	MsgStoreMemory
	MsgForward
	MsgShutdown
)

type Message struct {
	From    string
	To      string
	Type    MessageType
	Payload string
}
