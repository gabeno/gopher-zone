package main

import "fmt"

// fn takes two ints and returns an int
func plus(a int, b int) int {
	return a + b
}

// fn takes three ints and returns an int
// omit type name for like-typed params up to the last param that declares the type
func plusplus(a, b, c int) int {
	return a + b + c
}

// fn with multiple return values
func double() (int, int) {
	return 3, 7
}

// fn accepts a variable number of arguments
func sums(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	fmt.Println("1+2 =", plus(1, 2))
	fmt.Println("1+2+3 =", plusplus(1, 2, 3))
	fmt.Println("1+2+3 =", plusplus(1, 2, 3))

	a, b := double()
	fmt.Println("a:", a, "b:", b)

	sums(1, 2)
	sums(1, 2, 3)

	nums := []int{1, 2, 4, 5, 7}
	sums(nums...)
}
