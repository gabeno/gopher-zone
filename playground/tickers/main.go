// do something repeatedly at regular intervals
// tickers use a similar mechanism - a channel that is sent values

package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("tick at", t)
			}
		}
	}()

	// ticker can be stopped like timers after which it does not receive any more values
	// in its channel; this one stops after 1600ms
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("ticker stopped")
}
