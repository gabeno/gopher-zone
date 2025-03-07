// `defer` ensures that a function call is performed later in
// a program's execution usually for cleanup

package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/tmp/file.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func closeFile(f *os.File) {
	fmt.Println("closing file")
	err := f.Close()
	if err != nil {
		panic(err)
	}
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}
