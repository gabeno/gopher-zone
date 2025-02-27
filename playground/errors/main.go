// @todo: read blog https://go.dev/blog/go1.13-errors

package main

import (
	"errors"
	"fmt"
)

// by convention, errors are the last return value and have type `error` interface
func f(n int) (int, error) {
	if n == 42 {
		// `errors.New` constructs a basic error value with given message
		return -1, errors.New("can not work with 42")
	}
	// `nil` indicates there was no error
	return n + 3, nil
}

// sentinel error: a predeclared variable used to signify a specific error condition
var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(n int) error {
	if n == 2 {
		return ErrOutOfTea
	} else if n == 4 {
		// wrap errors with higher-level errors to add context using `%w` and `Errorf`
		// wrapped errors create a logical chain that can be queried with functions like
		// `errors.Is` and `errors.As`
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{40, 42} {
		// inline error check
		if r, err := f(i); err != nil {
			fmt.Println("f failed:", err)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("we should buy more tea")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("now it is dark!")
			} else {
				fmt.Println("unkown error %s\n", err)
			}
		}
		fmt.Println("Tea is ready!")
	}
}
