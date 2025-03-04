// timers represent a single event in the future;
// specify how long to wait, and it provides a channel that will
// be notified at that time

package main

import (
	"fmt"
	"time"
)

func main() {
	// timer waits 2 seconds
	timer1 := time.NewTimer(2 * time.Second)

	// blocks on timer's channel until it send a value indicating that
	// the timer fired
	<-timer1.C
	fmt.Println("timer 1 fired")

	// cancel the timer before it fires
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()
	stop := timer2.Stop()
	if stop {
		fmt.Println("timer 2 stopped")
	}

	// give timer2 enough time to fire if it was going to,
	// to show it is stopped
	time.Sleep(2 * time.Second)
}
