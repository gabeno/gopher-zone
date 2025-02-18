package main

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return factorial(n-1) * n
}

func main() {
	fmt.Println(factorial(7))

	// anonymous function can also be recursive
	// needs to be declared with a var to store
	// the function before its defined
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))

}
