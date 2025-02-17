package main

import "fmt"

func main() {
	// basic if/else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	fmt.Println("")

	// if statement without else
	if 8%2 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	fmt.Println("")

	// using logical operators && and ||
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 is divisible by 2")
	}
	fmt.Println("")

	// conditionals preceded by a statement
	if num := 9; num < 0 {
		fmt.Println("is negative")
	} else if num < 10 {
		fmt.Println("has single digit")
	} else {
		fmt.Println("has multiple digits")
	}
}
