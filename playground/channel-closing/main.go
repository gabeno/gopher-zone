package main

import "fmt"

func main() {
	// `jobs` channel communicates work to be done from main goroutine to
	// worker goroutine
	jobs := make(chan int, 5)
	done := make(chan bool)

	// worker goroutine receives jobs
	// `more` is false if `jobs` channel has been closed and all values received
	// we use this to notify on `done`
	go func() {
		for {
			for {
				i, more := <-jobs
				if more {
					fmt.Println("received job", i)
				} else {
					fmt.Println("received all jobs")
					done <- true
					return
				}
			}
		}
	}()

	// send 3 jobs to the channel then close it
	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("sent job", i)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// await worker using synchronization approach
	<-done

	_, ok := <-jobs
	fmt.Println("received more jobs", ok)
}
