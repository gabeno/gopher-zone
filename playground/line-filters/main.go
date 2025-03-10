// line filters are programs that read input on stdin, process it, and then
// prints some derived result to stdout eg `grep` and `sed`

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// wrap unbuffered `os.Stdin` with a buffered scanner giving a convenient
	// `Scan` method that advances the scanner to the next token which is the
	// next line in the default scanner
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// `Text` returns current token, ie next line, from the input
		ucl := strings.ToUpper(scanner.Text())
		// write out the uppercased line
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

//
//➜  playground git:(main) ✗ echo "hello" > /tmp/lines
//➜  playground git:(main) ✗ echo "filter" >> /tmp/lines
//➜  playground git:(main) ✗ cat /tmp/lines
//hello
//filter
//➜  playground git:(main) ✗ cat /tmp/lines | go run line-filters/main.go
//HELLO
//FILTER
//
