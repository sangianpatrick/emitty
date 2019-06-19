package emitty

// New will return event message
func New(debug bool) *Signal {
	ch := make(chan *Message)
	return &Signal{
		channel: ch,
		debug:   debug,
	}
}
