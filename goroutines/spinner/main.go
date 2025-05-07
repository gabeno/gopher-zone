package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond) // show a spinner as we compute fib result
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("fibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, c := range `-\|/` {
			fmt.Printf("\r%c", c)
			time.Sleep(delay)
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

/**
* when main function returns all goroutines are abruptly terminated an the program exits
* all goroutines return from main or exit the program
* there is no prgrammatic way for one goroutine to stop another but we can
* communicate with a goroutine to request it to stop itself
*
* this program has two anonymous activities - spinning and fibonacci computation
* each as a separate function but both running concurrently
 */
