package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	// const can be declared anywhere a var is
	const n = 500000000

	// const performs arithmetic with constant precision
	const d = 3e20 / n
	fmt.Println(d)

	// numeric constant has no type until it is given one
	fmt.Println(int64(d))

	// numeric const can assume any type based on usage context
	// eg Sin expects a float64
	fmt.Println(math.Sin(n))
}
