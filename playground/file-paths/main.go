// `filepath` provides functions to parse and construct file paths
// in a way that is portable across different OSes

package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// `Join` to construct paths in a portable way
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// removes superfluous separators and directory changes
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// split path to directory and file
	fmt.Println("Dir:", filepath.Dir(p))
	fmt.Println("Base:", filepath.Base(p))

	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// find filename without the extension
	fmt.Println(strings.TrimSuffix(filename, ext))

	// find the relative path between a base and target or return an error
	// if target can not be made relative to base
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
