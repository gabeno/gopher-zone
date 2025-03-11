package main

import (
	"fmt"
	"net/http"
	"time"
)

// fundamental concept of servers is handlers
// a handler is an object implementing `http.Handler` interface
// a common way is to use `http.HandlerFunc` adapter on functions with appropriate signature

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// register handlers on server routes using `http.HandleFunc` convenience function
	// which sets up a default router
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/headers", headers)

	// listen and serve requests
	// `nil` means we use default router we just set up.
	http.ListenAndServe(":8080", nil)
}

/*
➜  playground git:(main) ✗ go run http-server/main.go &
[1] 86711
➜  playground git:(main) ✗ curl localhost:8080/headers
User-Agent: curl/8.7.1
Accept: *\/*
➜  playground git:(main) ✗ curl localhost:8080/greet
Hello World! 2025-03-11 10:44:13.26458 +0300 EAT m=+7.457958501%
*/
