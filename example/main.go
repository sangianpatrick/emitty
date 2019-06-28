package main

import (
	"fmt"
	"time"

	"github.com/sangianpatrick/emitty"
)

var debug bool

func init() {
	debug = true
}

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Panic Recovered", r)
		}
	}()

	fmt.Println("Emitty Simple Usage")

	signal := emitty.New(debug)
	listener := emitty.NewListener(&emitty.Config{
		Signal:          signal,
		NumberOfWorkers: 3,
	})
	emitter := emitty.NewEmitter(signal)

	err := listener.AttachEvent(&emitty.Event{
		Name:             "printStr",
		ActiveOn:         time.Now().Add(time.Second * 0),
		Expiration:       time.Second * 15,
		Handler:          exampleHandler,
		MaxHits:          5,
		StartImmediately: true,
	})

	if err != nil {
		panic(err)
	}

	listener.Start()

	time.Sleep(time.Second * 3)

	emitter.Emit("printStr", "Hello World\n")

	fmt.Scanln()
}

func exampleHandler(args ...interface{}) {
	if str, ok := args[0].(string); ok {
		fmt.Printf("String: %s", str)
	}
}
