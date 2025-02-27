// goroutine is a lightweight thread of execution

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
	// run `f` synchronously
	f("direct")

	// executed concurrently
	go f("goroutine")

	// go routine for anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// adds a delay for goroutine functions to run and finish
	time.Sleep(time.Second)
	fmt.Println("done")
}

// output of blocking call printed first
// then output of two go routines which may be interleaved
// because they are being called concurrently by Go runtime
