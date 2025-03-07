// different string formating options

package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	// print an instance of point struct
	fmt.Printf("struct1: %v\n", p)

	// %+v includes structs fields names if value is a struct
	fmt.Printf("struct2: %+v\n", p)

	// print syntax representation of the value
	fmt.Printf("struct3: %#v\n", p)

	// print type of a value
	fmt.Printf("type: %T\n", p)

	// formatting booleans
	fmt.Printf("bool: %t\n", true)

	// format integer with base 10 representation
	fmt.Printf("integer: %d\n", 13)

	// print binary representation
	fmt.Printf("binary: %b\n", 14)

	// print character corresponding to the given integer
	fmt.Printf("char: %c\n", 33)

	// provide hex encoding
	fmt.Printf("hex: %x\n", 143)

	// basic float formatting with %f
	fmt.Printf("float1: %f\n", 1.2)

	// formart floats in different scientific representations
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	// basic string printing with %s
	fmt.Printf("str1: %s\n", "\"string\"")

	// double quote strings as in Go source code use %q
	fmt.Printf("str2: %q\n", "\"string\"")

	// hex renders string in base 16 with two output characters per byte of input
	fmt.Printf("str3: %x\n", "hex this")

	// print representation of a pointer
	fmt.Printf("ptr: %p\n", &p)

	// width and precision when printing numbers
	// by default right justified and padded with spaces
	// left justify with `-` flag
	fmt.Printf("width1: |%6d|%-6d|%6d|\n", 12, 345, 1024)

	// width and precision for floats
	// specify width and decimal precision at once
	fmt.Printf("width2: |%6.2f|%6.2f|%6.2f|\n", 1.2, 34.5, 1.024)

	// to left justify use `-` flag
	fmt.Printf("width3: |%-6.2f|%-6.2f|%-6.2f|\n", 1.2, 34.5, 1.024)

	// width when printing strings
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// left justify similar to numbers
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// `SPrintf` formats and returns a string without printing it anywhere
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// formart and print to `io.Writer`s other than `osStdout` using `FPrintf`
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
