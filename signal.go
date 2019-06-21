package emitty

// Signal contains channel that bridging the Emitter and Listener to have communication.
type Signal struct {
	channel chan *Message
	debug   bool
}

// Close will release the communication between Emitter and Listener
func (s *Signal) Close() {
	close(s.channel)
}
