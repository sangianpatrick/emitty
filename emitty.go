package emitty

// New will return event message
func New() chan *Message {
	ch := make(chan *Message)
	return ch
}
