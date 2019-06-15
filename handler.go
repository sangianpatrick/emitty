package emitty

// Handler is a type for function that will be registred and executed
// when the listener catch an event.
type Handler func(...interface{})
