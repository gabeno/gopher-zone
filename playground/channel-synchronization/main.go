// we can use channels to syncronize goroutine workers
// WaitGroup is a better alternative for this

package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working ...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// notify another goroutine that this function's work is done
	done <- true
}

func main() {
	done := make(chan bool)
	// start a worker goroutine giving it the channel to notify on
	go worker(done)

	// block until we receive a notification from the worker on this channel
	<-done
}
