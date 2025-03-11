// simple client to issue http requests

package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	// `http.Get` is a convenient shortcut for creating an `http.Client` object
	// and calling its `Get` method; it uses `http.DefaultClient` object which
	// has useful default settings
	res, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("Response status:", res.Status)

	// print the first 5 lines of the response body
	scanner := bufio.NewScanner(res.Body)
	for i := 0; i < 5 && scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
