// managing state via atomic counters

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// atomic integer type to represent our counter
	var ops atomic.Uint64

	// wait for all goroutines to finish their work
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		// launch 50 goroutines each incrementing counter exactly 1000 times
		go func() {
			for i := 0; i < 1000; i++ {
				// atomically increase counter
				ops.Add(1)
			}
			wg.Done()
		}()
	}

	// wait till all goroutines are done
	wg.Wait()

	// `Load` safely reads a value while other goroutines are updating it
	// prints exactly 50000 operations
	fmt.Println("ops", ops.Load())
}

// had we used a non-atomic integer and incremented it with `ops++`, we'd likely get a different
// number, changing between runs, because the goroutines would interfere with each other.
// also we would get data race failures when running with `-race`
