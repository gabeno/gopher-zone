// by default, channels are unbuffered, that is
// they only accept a send (chan <-) if there is a
// corresponding receive (<- chan)
// Buffered channels accept a limited number of values
// without a corresponding receiver for those values.

package main

import "fmt"

func main() {
	// channel of strings buffering up to 2 values
	messages := make(chan string, 2)

	// because this channel is buffered, we can send these two values
	// without corresponding concurrent receive
	messages <- "buffered"
	messages <- "channel"

	// later receive values as usual
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
