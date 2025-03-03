// timeouts are important for programs that connect
// to external resources or need to bound execution time
// we use channels and select

package main

import (
	"fmt"
	"time"
)

func main() {

	// simulate external call that runs for 2s and then sends message to channel
	// channel is buffered so send in the goroutine is non blocking - a pattern
	// to prevent goroutine leaks in case the channel is never read

	// set shorter timeout => prints "timeout 1"
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case msg := <-c1:
		fmt.Println(msg)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	// set longer timeout => prints "result 2"
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case msg := <-c2:
		fmt.Println(msg)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 1")
	}
}
