package main

import (
	"fmt"
	"time"

	"github.com/sangianpatrick/emitty"
)

func main() {
	defer func() {
		recover()
		fmt.Println("Program Exit")
	}()
	signal := emitty.New(true)
	listener := emitty.NewListener(&emitty.Config{
		Signal:          signal,
		NumberOfWorkers: 3,
	})

	err := listener.AttachEvent(&emitty.Event{
		Name:             "printStr",
		ActiveOn:         time.Now().Add(time.Second * 0),
		Expiration:       time.Second * 60,
		Handler:          exampleHandler,
		MaxHits:          5,
		StartImmediately: true,
	})

	listener.Start()

	emitter := emitty.NewEmitter(signal)

	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 5)

	emitter.Emit("printStr", "Patrick")

	fmt.Scanln()
}

func exampleHandler(data ...interface{}) {
	if str, ok := data[0].(string); ok {
		fmt.Printf("String: %s", str)
	}
}
