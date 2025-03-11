// Enable Go program to intelligently handle Unix Signals eg SIGTERM
// we leverage Go channels

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// send os.Signal values on a buffered channel
	// we create a channel to receive these notifications
	sigs := make(chan os.Signal, 1)

	// register channel to receive notifications of the specified signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// we can receive from `sigs` in main but showing how it works in a separate
	// goroutine to demo a more realistic scenario of graceful shutdown
	done := make(chan bool, 1)

	// goroutine executes a blocking receive for signals
	// when done it will print out and then notify the program that it can
	// finish
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// program will wait here till it gets expected signal (goroutine sending a
	// value on `done`) and then exit
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}

/*
*
➜  playground git:(main) ✗ go run signals/main.go
awaiting signal
^C
interrupt
exiting
* */
