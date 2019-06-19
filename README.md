# emitty

A simple Event Emitter package for Golang application. This package contains both Emitter and Event.
Emitter only has one function "Emit()" that send the data to the listener. Before the listener do its jobs, Event should attach an event name and a handler (to be executed after event is catched) by AttachEvent function, but if the event is needless, it could be detached by DetachEvent function.

## How to install

Using go get :

```go get -u github.com/sangianpatrick/emitty```

Using DEP :

```dep ensure -add github.com/sangianpatrick/emitty```

## How to use

This is the package implementation.

```
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
```
