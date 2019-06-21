package emitty

import (
	"fmt"
	"time"
)

// Listener acts as Signal's observer and contains behavior to attach and detach an event
type Listener interface {
	Start()
	AttachEvent(event *Event) error
	DetachEvent(eventName string) error
}

type listOfEvent map[string]*Event

type listener struct {
	c   *Config
	loe listOfEvent
}

// NewListener will prepare the workers and return function to start, attach & detach event
func NewListener(c *Config) Listener {
	l := &listener{
		c:   c,
		loe: make(listOfEvent),
	}

	return l
}

// Start will run the listener.
func (l *listener) Start() {
	if l.c.NumberOfWorkers <= 1 {
		l.doListen()
		l.c.Signal.log("Listener", "Running on 1 worker", nil, nil)
		return
	}
	for i := 1; i <= int(l.c.NumberOfWorkers); i++ {
		l.doListen()
	}
	l.c.Signal.log("Listener", fmt.Sprintf("Running on %v workers", l.c.NumberOfWorkers), nil, nil)
}

// AttachEvent will push the event to the listener
// and label it with its name.
func (l *listener) AttachEvent(e *Event) error {
	err := l.eventAttachmentScheduler(e)
	return err
}

// DetachEvent will remove the event from list in the listener.
func (l *listener) DetachEvent(eventName string) error {
	if _, ok := l.loe[eventName]; ok {
		l.doDetach(eventName)
		return nil
	}
	return ErrDetach
}

func (l *listener) eventAttachmentScheduler(e *Event) error {
	unixNow := int(time.Now().Unix())
	aoUnixTime := int(e.ActiveOn.Unix())
	timeToActivate := aoUnixTime - unixNow

	if e.StartImmediately {
		l.doAttach(e)
		return nil
	}

	if timeToActivate < 1 {
		return ErrAttach
	}

	setSchedule := func() {
		go func() {
			select {
			case <-time.After(time.Duration(timeToActivate) * time.Second):
				l.doAttach(e)
			}
		}()
	}

	setSchedule()

	return nil
}

func (l *listener) doAttach(e *Event) {
	go func() {
		l.loe[e.Name] = e
		l.eventDetachmentScheduler(e)
	}()
}

func (l *listener) eventDetachmentScheduler(e *Event) {
	setSchedule := func() {
		go func() {
			select {
			case <-time.After(e.Expiration):
				l.doDetach(e.Name)
			}
		}()
	}

	if int(e.Expiration) > 0 {
		setSchedule()
	}

}

func (l *listener) doDetach(eventName string) {
	go func() {
		delete(l.loe, eventName)
		l.c.Signal.log("Listener", fmt.Sprintf("Event with name '%s' has been detached", eventName), nil, nil)
	}()
}

func (l *listener) doListen() {
	go func() {
		for {
			select {
			case m := <-l.c.Signal.channel:
				if e, ok := l.loe[m.EventName]; ok {
					l.execute(e.Handler, m.EventName, m.Data)
					continue
				}
				l.c.Signal.log("Listener", fmt.Sprintf("On Event: '%s'", m.EventName), nil, errEventNotFound)
			}
		}
	}()
}

func (l *listener) execute(fn Handler, eventName string, data []interface{}) {
	go fn(data...)
	l.c.Signal.log("Listener", fmt.Sprintf("Executing handler on event '%s'", eventName), nil, nil)
}
