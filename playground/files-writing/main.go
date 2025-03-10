package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// dump a string (or bytes) into a file
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// for more granular writes, open a file
	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()

	// write byte slices
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// WriteString is also available
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// flush writes to stable storage
	f.Sync()

	// bufio provides buffered writers in addition to the buffered readers
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// flush to ensure all buffered operations have been applied to
	// the underlying writer
	w.Flush()
}
