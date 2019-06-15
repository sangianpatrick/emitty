package emitty

import (
	"log"
)

// Event contains attach and detach handler.
type Event interface {
	AttachEvent(eventName string, handler Handler) error
	DetachEvent(eventName string) error
}

type listofHandler map[string]Handler

type event struct {
	ch  <-chan *Message
	loh listofHandler
}

// NewEvent is an constructor of event.
func NewEvent(ch <-chan *Message) Event {
	loh := make(listofHandler)
	e := &event{
		ch:  ch,
		loh: loh,
	}

	e.doListen()

	return e
}

func (e *event) doListen() {
	go func() {
		for {
			select {
			case m := <-e.ch:
				if fn, ok := e.loh[m.EventName]; ok {
					e.execute(fn, m.Data)
					continue
				}
				log.Printf("Event name: '%s', %s", m.EventName, errEventNotFound)
			}
		}
	}()
}

func (e *event) execute(fn Handler, data []interface{}) {
	go fn(data...)
}

func (e *event) attach(en string, h Handler) {
	e.loh[en] = h
}

func (e *event) detach(en string) {
	delete(e.loh, en)
}

// AttachEvent is a function that will push the event handler to the listener
// and label it with event name.
func (e *event) AttachEvent(eventName string, handler Handler) error {
	e.attach(eventName, handler)
	if _, ok := e.loh[eventName]; !ok {
		return ErrAttach
	}
	return nil
}

// DetachEvent is a function that will remove the event handler from the listener
// by its label.
func (e *event) DetachEvent(eventName string) error {
	if _, ok := e.loh[eventName]; ok {
		e.detach(eventName)
		return nil
	}
	return ErrDetach
}
