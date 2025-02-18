package main

import "fmt"

// return another function which is anonymously defined in intSeq
// the returned function closes over variable i
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// newInt is a function assigned a value of i and incremented
	// on every call
	newInt := intSeq()
	fmt.Println(newInt())
	fmt.Println(newInt())
	fmt.Println(newInt())

	// new state with another call to intSeq
	anotherInt := intSeq()
	fmt.Println(anotherInt())
}
