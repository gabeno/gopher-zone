package main

import "fmt"

func main() {
	// declare a variable with `var` keyword
	var a = "inital"
	fmt.Println(a)

	// declare multiple variables
	var x, c int = 1, 2
	fmt.Println(x, c)

	// non-initialized variables are zero-valued
	var d int
	fmt.Println(d)

	// := syntax for declaring and initializing a variable
	e := "apple"
	fmt.Println(e)

	// variable names can have unicode characters besides upper and lowercase and underscore
	手紙 := "letter"
	fmt.Println(手紙)

	fmt.Println()

	// variables without intial values are set to their corresponding zero values
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("initial values -> int:%v float:%v bool:%v string:%q\n", i, f, b, s)
}
