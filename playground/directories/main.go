package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// create new subdirectory in current working directory
	err := os.Mkdir("subdir", 0755)
	check(err)

	// `os.RemoveAll` deletes a whole directory tree; similar to `rm -rf`
	defer os.RemoveAll("subdir")

	// helper
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// create hierarchy of directories; similar to `mkdir -p`
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	c, err := os.ReadDir("subdir/parent")
	check(err)

	// list subdir/parent
	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// change dir
	err = os.Chdir("subdir/parent/child")
	check(err)

	// see all contents of subdir/parent/child when listing current directory
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// cd back to where we started
	err = os.Chdir("../../..")
	check(err)

	// visit a directory
	// callback handles every file or directory visited
	fmt.Println("visiting subdir")
	err = filepath.WalkDir("subdir", visit)
	check(err)
}

func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
