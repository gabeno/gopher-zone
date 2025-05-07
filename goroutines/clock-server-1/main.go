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
		handleConn(conn) // handle one connection at a time
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

// two clients are each handled sequentially
// the second client must wait until the first client is finished
// because the server is sequential

/**
*
->  goroutines git:(main) $ go run clock-server-1/main.go&
[1] 67165
->  goroutines git:(main) $ nc localhost 8000
10:53:38
10:53:39
10:53:40
10:53:41
10:53:42
10:53:43
10:53:44
10:53:45
10:53:46
^C
*
*/

/**
*
➜  ~ nc localhost 8000
10:53:48
10:53:49
10:53:50
10:53:51
10:53:52
^C
*/
