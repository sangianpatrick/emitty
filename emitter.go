package emitty

// Emitter contains function to emit then event
type Emitter interface {
	Emit(eventName string, data ...interface{})
}

type emitter struct {
	ch chan<- *Message
}

// NewEmitter will return emitter
func NewEmitter(ch chan<- *Message) Emitter {
	return &emitter{
		ch: ch,
	}
}

// Emit is function that will send the data to the listener
func (e *emitter) Emit(eventName string, data ...interface{}) {
	message := &Message{
		EventName: eventName,
		Data:      data,
	}
	e.ch <- message
}
