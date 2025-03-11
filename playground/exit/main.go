// exit with a given status

package main

import (
	"fmt"
	"os"
)

func main() {
	// will never be run when using `os.Exit` so "!" is never printed
	defer fmt.Println("!")

	os.Exit(3)
}
