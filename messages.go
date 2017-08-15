// Package writer provides message writer concept definitions.
package writer

// Message represents a message emitted by the mainflux adapters layer.
type Message struct {
	Channel     string
	ContentType string
	Publisher   string
	Protocol    string
	Payload     []byte
}

// MessageRepository specifies a message persistence API.
type MessageRepository interface {
	// Save persists the message. A non-nil error is returned to indicate
	// operation failure.
	Save(Message) error
}
