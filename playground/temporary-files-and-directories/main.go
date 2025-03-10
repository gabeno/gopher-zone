// Temporary fileas and directories are useful for keeping data that
// is not needed after program exits; and they do not pollute the
// filesystem

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// create file at default location due to "" and
	// open it for reading and writing at default locationi
	f, err := os.CreateTemp("", "sample")
	check(err)

	fmt.Println("temp file name:", f.Name())

	defer os.Remove(f.Name())

	// write data to the file
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// create a temp directory
	dname, err := os.MkdirTemp("", "sampleDir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
