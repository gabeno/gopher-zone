package main

import "fmt"

func main() {
	// declare a variable
	var a = "inital"
	fmt.Println(a)

	// declare multiple variables
	var b, c int = 1, 2
	fmt.Println(b, c)

	// non-initialized variables are zero-valued
	var d int
	fmt.Println(d)

	// := syntax for declaring and initializing a variable
	e := "apple"
	fmt.Println(e)
}
