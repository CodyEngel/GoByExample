package main

import "fmt"

func main() {
	// Here we make a channel of strings buffering up to 2 values.
	messages := make(chan string, 2)

	// Because this channel is buffered, we can send these values into the channel without a corresponding receive.
	// If the channel was not buffered, this the following code would terminate with a deadlock. Similarly, if we write
	// more than 2 values to the buffer we will get a deadlock as well.
	messages <- "buffered"
	messages <- "channel"

	// Later we can receive these two values as usual.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
