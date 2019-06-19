package emitty

// Emitter contains function to emit then event
type Emitter interface {
	Emit(eventName string, data ...interface{})
}

type emitter struct {
	signal *Signal
}

// NewEmitter will return emitter
func NewEmitter(s *Signal) Emitter {
	return &emitter{
		signal: s,
	}
}

// Emit is function that will send the data to the listener
func (e *emitter) Emit(eventName string, data ...interface{}) {
	message := &Message{
		EventName: eventName,
		Data:      data,
	}
	e.signal.channel <- message
}
