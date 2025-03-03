// channels are pipes that connect goroutines,
// you can send values into channels from one goroutine
// and receive those values into another goroutine.

package main

import "fmt"

func main() {
	// create a new channel with `make` typed as `string` - channel conveys strings
	messages := make(chan string)

	// send a value into a channel with `channel <-` syntax
	// we send the message inside a goroutine
	go func() {
		messages <- "ping"
	}()

	// receive a value with `<- channel` syntax
	msg := <-messages

	fmt.Println(msg)
}
