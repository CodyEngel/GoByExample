package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Synchronous call to function f.
	f("direct")

	// To invoke this function in a goroutine, use go f(s). This will execute the new goroutine concurrently.
	go f("goroutine")

	// You can also start a goroutine from an anonymous function call.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two functions are running asynchronously in separate goroutines. We'll wait for them to finish with a simple
	// thread sleep. We could also make use of WaitGroup which would be the preferred option.

	time.Sleep(time.Second)
	fmt.Println("done")
}