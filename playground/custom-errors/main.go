// possible to use custom types as `errors` by implementing `Error()` on them

package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg     int
	message string
}

func (ae *argError) Error() string {
	return fmt.Sprintf("%d - %s", ae.arg, ae.message)
}

func f(n int) (int, error) {
	if n == 42 {
		return -1, &argError{n, "can not work with 42"}
	}
	return n + 3, nil
}

func main() {
	_, err := f(42)
	var ae *argError
	// `errors.As` checks that a given error (or an error in its chain) matches a specific
	// error type and converts to a value of that type returning true. if there is no match,
	// it returns false
	if errors.As(err, &ae) {
		fmt.Println(ae.arg, ae.message)
	} else {
		fmt.Println("error does not match argError")
	}
}
