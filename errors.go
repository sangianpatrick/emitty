package emitty

import "errors"

// ErrAttach it will be returned if there is error detected while attaching event handler.
var ErrAttach = errors.New("Cannot attach event handler")

// ErrDetach it will be returned if there is error detected while detaching event handler.
var ErrDetach = errors.New("Cannot detach event handler")

// ErrEventNotFound
var errEventNotFound = errors.New("Event is not found")
