package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	// arg passed by value
	// zeroval gets a copy of ival
	zeroval(i)
	fmt.Println("zeroval:", i)

	// arg passed by reference
	// zero pointer takes an int pointer
	// *iptr dereferences the pointer from its memory address to the current value at that address
	// assigned value to a dereferences pointer changes the value at the references address
	// &i gives the memory address of i ie pointer to i
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// print pointer
	fmt.Println("pointer:", &i)
}
