package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// reading files requires most calls for errors
// helper to streanline error checks
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// lets use an in memory file
	data, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Println(string(data))

	f, err := os.Open("/tmp/dat")
	check(err)
	defer f.Close()

	// read some bytes from the beginning of the file, allowing up to 5 bytes
	// but also note how many actually were read
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// seek to a known location in the file and read from there
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// seek relative to current cursor position
	_, err = f.Seek(2, io.SeekCurrent)
	check(err)

	// seek relative to end of file
	_, err = f.Seek(-4, io.SeekEnd)
	check(err)

	// `io` package provides some functions that may be helpful for file reading
	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// there is no builtin rewind but below accomplishes this
	_, err = f.Seek(0, io.SeekStart)
	check(err)

	// reading with bufio
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
}
