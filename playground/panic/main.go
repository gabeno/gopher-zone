// `panic` means something went unexpectedly wrong
// fail fast on errors that should not occur or that we are not
// prepared to handle gracefully

package main

import "os"

func main() {
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

	panic("a problem here")
}
