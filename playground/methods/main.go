package main

import "fmt"

type rectangle struct {
	width, height int
}

// pointer receiver
func (r *rectangle) area() int {
	return r.height * r.width
}

// value receiver
func (r rectangle) perimerer() int {
	return 2*r.height + 2*r.width
}

func main() {
	r := rectangle{width: 27, height: 43}

	fmt.Println("area: ", r.area())
	fmt.Println("perimeter: ", r.perimerer())

	rPtr := &r

	// automatically handle conversion between values and pointers for method calls
	// use pointer receiver type to avoid:
	// - copying on method calls
	// - the method to mutate the receiviing struct
	fmt.Println("area: ", rPtr.area())
	fmt.Println("perimeter: ", rPtr.perimerer())
}
