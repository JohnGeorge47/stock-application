package socket

//Message is the message type for the ws
type Message struct {
	MessageID string
	UserID    string
	Data      []byte
}
