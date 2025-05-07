package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn) // handle multiple connections
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		// handle echos concurrently allowing them to overlap in time within a single connection
		go echo(c, input.Text(), 1*time.Second)
	}
}

func echo(c net.Conn, text string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(text))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", text)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(text))
	time.Sleep(delay)
}
