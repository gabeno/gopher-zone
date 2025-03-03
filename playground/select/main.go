// `select` lets you wait on multiple channel operations
// goroutines + channels is powerful

package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// each channel will receive a value after some amount of time
	// eg blocking RPC in concurrent goroutines
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	// use select to await each of these values simultaneously, printing
	// each one as it arrives
	// total exection time ~2s since sleeps are concurrent
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)

		}
	}
}
