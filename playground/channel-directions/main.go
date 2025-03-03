// when using channels as function parameters,
// you can specify if a channel is meant to only
// send or receive values - better type-safety

package main

import "fmt"

// param is a channel for sending values
func ping(sendTo chan<- string, msg string) {
	sendTo <- msg
}

// function accepts one channel for receives and other for sends
func pong(readFrom <-chan string, sendTo chan<- string) {
	msg := <-readFrom
	sendTo <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message") // write to pings
	pong(pings, pongs)            // read from pings and write to pongs
	fmt.Println(<-pongs)
}
