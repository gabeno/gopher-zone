// `context.Context` controls cancellation.
// it carries deadlines, cancellation signals and other request-scoped values
// across API boundaries and goroutines

package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	// context is created by net/http and is available with `Context()` method
	ctx := r.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	// wait a few seconds before sending a reply to the client
	// keep an eye on context's `Done()` for a signal that we should cancel the work
	// and return as soon as possible
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// error returned explains why the `Done()` channel was closed
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8080", nil)
}

/*
➜  playground git:(main) ✗ go run http-context/main.go&
[1] 87890
➜  playground git:(main) ✗ curl localhost:8080/hello
server: hello handler started
server: hello handler ended
hello
➜  playground git:(main) ✗ curl localhost:8080/hello
server: hello handler started
^Cserver: context canceled
server: hello handler ended
*/
