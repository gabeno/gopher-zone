// basic sends and recevies on channels are blocking
// select with a default achieves non-blocking sends,
// receives and non-blocking multiway selects

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// non-blocking receive
	// if value is available on `messages` channel, it is printed
	// else default case is selected
	// no message sent so default case applies
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// non-blocking receive
	// `msg` can not be sent to `messages` channel because channel has no buffer and there is no receiver
	// default case selected
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// multi-way non-blocking select
	// attempt non-blocking receives on both `messages` and `signals`
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
