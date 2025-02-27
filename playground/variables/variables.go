package main

import "fmt"

// package scope only allows variable declaration and not the short form
var global int

// global := 4 // <= this is not alloed

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

	// := syntax for declaring and initializing a variable -> short declaration
	// variable values already known
	e := "apple"
	fmt.Println(e)

	// we can redeclare existing variables as we create new ones
	e, k := "new", 100
	fmt.Println(e, k)

	// variable names can have unicode characters besides upper and lowercase and underscore
	手紙 := "letter"
	fmt.Println(手紙)

	fmt.Println()

	// variables without intial values are set to their corresponding zero values
	// used when we dont know initial values
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("initial values -> int:%v float:%v bool:%v string:%q\n", i, f, b, s)

	// group related variables together for better readability
	var (
		// related
		video string

		// closely related
		current  int
		duration int
	)
	fmt.Printf("%q, %v, %v\n", video, current, duration)

	// type conversion
	// create a new values changing from one type to another
	speed := 100 // int
	force := 2.5 // float64

	// => 200, lost precision due to rounding down during int to float conversion
	fmt.Println("speed * force =", speed*int(force))
	fmt.Println(force, int(force))

	// => 250,
	fmt.Println("speed * force =", float64(speed)*force)
}
