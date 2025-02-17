package main

import (
	"fmt"
)

func main() {
	// for loop with single condition
	i := 0
	for i < 3 {
		fmt.Println(i)
		i += 1
	}
	fmt.Println("")

	// do something n times with classic for loop
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("")

	// do something n times with range
	for r := range 3 {
		fmt.Println(r)
	}
	fmt.Println("")

	// for without condition will run repeatedly unless
	// break and return from enclosing function
	for {
		fmt.Println("loop")
		break
	}
	fmt.Println("")

	// continue to next iteration of the loop
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
