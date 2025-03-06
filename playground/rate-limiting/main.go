// rate limiting controls resource utilization and quality of service
// implement using goroutines, channels and tickers

package main

import (
	"fmt"
	"time"
)

func main() {
	// create a channel to mimic incoming requests
	// we will serve these off a channel with the same name
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// limiter channel receives a value every 200ms
	// this is the regulator in our rate limiter scheme
	limiter := time.Tick(200 * time.Millisecond)

	// block on a receive from limiter channel before serving each request
	// we limit ourselves to 1 request every 200ms
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// buffer our limiter channel with 3 events
	burstyLimiter := make(chan time.Time, 3)

	// fill up the channel to simulate allowed bursting
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// every 200ms try add a new value to burstyLimiter upto its limit
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// simulate 5 more incoming requests, the first 3 will benefit from the
	// bursty capability of burstyLimiter
	burstRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstRequests <- i
	}
	close(burstRequests)

	for req := range burstRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
