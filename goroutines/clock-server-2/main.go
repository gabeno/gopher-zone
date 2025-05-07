// clock server that periodically writes time to the client

package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// listen for incoming connection on tcp localhost port 8000
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		// blocks until a connection request is made
		// returns net.Conn object that represents the connection
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	// write current time to client from a loop
	// when client disconnects, writes fail and we return and close connection
	// handleConn continues to wait for next connection
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

// two clients are handled concurrently
// the second client does not wait until the first client is finished
// because the server is concurrent
// handleConn is a goroutine

/**
*
->  goroutines git:(main) $ go run clock-server-1/main.go&
[1] 67165
->  goroutines git:(main) $ nc localhost 8000
11:03:14
11:03:15
11:03:16
11:03:17
11:03:18
11:03:19
11:03:20
11:03:21
11:03:22
11:03:23
^C
*/

/**
*
->  ~ nc localhost 8000
11:03:17
11:03:18
11:03:19
11:03:20
11:03:21
11:03:22
11:03:23
11:03:24
11:03:25
^C
*/
