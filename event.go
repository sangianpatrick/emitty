package emitty

import (
	"time"
)

// Event contains property of event.
// Name represents the name of event, it will be used on emitter for sending data to the right place.
// Handler is a function that runs after the event is emitted.
// Expiration is set to tells the listener when to detach the event, if set to 0 (zero),
// the event will be active and ready anytime.
// ActiveOn will allow the listener to set a schedule when to activate the event,
// it must be set to at least 1 (one) minutes from now,
// or it will return an error, so if the needs is to activate immediately, set the StartImmediately to true
// and ActiveOn will be ignored by the listener.
// MaxHits can control the limit of access to the event, after it is reached, the event will be detached automatically,
// so set it to 0 (zero) for unlimited access.
type Event struct {
	Name             string
	Handler          Handler
	Expiration       time.Duration
	ActiveOn         time.Time
	StartImmediately bool
	MaxHits          int
}
