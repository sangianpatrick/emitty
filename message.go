package emitty

// Message contains event name and the data
type Message struct {
	EventName string
	Data      []interface{}
}
