package main

import "fmt"

func main() {
	// iterate over two values in the queue channel
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// iterate on each element as its received from queue
	// iteration terminates after printing all values since we closed the channel
	for el := range queue {
		fmt.Println(el)
	}
}
