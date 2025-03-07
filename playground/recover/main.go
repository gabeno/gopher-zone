// `recover` can stop a `panic` from aborting the program
// and let it continue with execution
// eg a server should not crash if one of the clients exhibits a critical error,
// instead close the connection and continue serving other clients

package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {
	// recover must be called from within a defer function
	// when the enclosing function panics, defer will activate and a recover
	// call within it will catch the panic

	defer func() {
		// return value of recover is the error raised in the call to panic
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// this code will not run because execution of main stops at the point of
	// panic and resumes in the deferred closure
	fmt.Println("after mayPanic()")
}
