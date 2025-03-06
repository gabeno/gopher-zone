// safely access data across multiple goroutines

package main

import (
	"fmt"
	"sync"
)

// use a mutex to synchronize access
// mutexes must not be copied
// if this struct is passed around it should be done by pointers
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	// zero value of a mutex is usable as-is so no initialization required here
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup

	// increment a named counter in a loop
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	// run several goroutines concurrently accessing same container and counter
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	// wait for goroutines to finish
	wg.Wait()
	fmt.Println(c.counters)
}
