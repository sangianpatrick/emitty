# emitty

A simple Event Emitter package for Golang application

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
	M := emitty.New()

	event := emitty.NewEvent(M)
	emitter := emitty.NewEmitter(M)

	eventName := "print"

	fmt.Printf("Attaching event with name '%s' ...\n", eventName)
	event.AttachEvent(eventName, printSomethingHandler)

	time.Sleep(time.Second)

	fmt.Printf("Emitting an event with name '%s' ...\n", eventName)

	time.Sleep(time.Second)

	emitter.Emit(eventName, "Hello World\n")

	time.Sleep(time.Second)

	fmt.Printf("Detaching event with name '%s' ...\n", eventName)

	time.Sleep(time.Second)

	event.DetachEvent(eventName)

	time.Sleep(time.Second)

	fmt.Printf("Emitting event  with name '%s' ...\n", eventName)

	time.Sleep(time.Second)

	emitter.Emit(eventName, "The event is still exist\n")

	fmt.Scanln()

}

func printSomethingHandler(data ...interface{}) {
	for _, v := range data {
		if str, ok := v.(string); ok {
			fmt.Printf("String: %s", str)
		}
	}
}
```
