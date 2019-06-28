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
```

The result will be:

```
Emitty Simple Usage
Emitty [INFO]: Name: Listener | Message: Running on 3 workers | Data <nil>
Emitty [INFO]: Name: Listener | Message: Executing handler on event 'printStr' | Data <nil>
String: Hello World
Emitty [INFO]: Name: Listener | Message: Event with name 'printStr' has been detached | Data <nil>
```
