// wait for multiple goroutines to finish

package main

import (
	"fmt"
	"sync"
	"time"
)

// function to run in every goroutine
func worker(id int) {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func main() {
	// wait for all goroutines lanched here to finish
	// if WaitGroup is explicitly passed into functions, it should be done by pointer
	var wg sync.WaitGroup

	// launch several goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// wrap the worker call in a closure that tells the WaitGroup that this worker
		// is done. this way the worker itself does not have to be aware of the
		// concurrency primitives involved in its execution
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// block until the WaitGroup counter goes back to 0, all the workers notified
	// they are done
	// this approach does not propagate errors from workers
	// @todo - see https://pkg.go.dev/golang.org/x/sync/errgroup for more advanced use cases
	wg.Wait()
}
