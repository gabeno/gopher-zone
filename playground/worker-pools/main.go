// implement worker pools using goroutines and channels

package main

import (
	"fmt"
	"time"
)

// worker for which we will run several concurrent instances
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker id", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker id", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	// make two channels to send work and collect results
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// starts 3 workers, initially blocked because there are no jobs yet
	for w := 0; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// send 5 jobs and then close
	for j := 0; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// collect all results of the work
	// also ensures worker goroutines have finished
	for a := 0; a <= numJobs; a++ {
		<-results
	}
}
