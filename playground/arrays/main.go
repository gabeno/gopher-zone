// arrays: sequence of items of a specific length

package main

import "fmt"

func main() {
	// array of 5 ints
	var a [5]int
	fmt.Println("emp: ", a)

	// set a value
	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])
	fmt.Println("len: ", len(a))

	// initialize and declare array
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl1: ", b)

	// let compiler determine number of elements
	c := [...]int{2, 4, 6, 8, 10}
	fmt.Println("dcl2: ", c)

	// specify the index with :, elements in between will be zeroed
	d := [...]int{1, 4: 4, 5}
	fmt.Println("decl3: ", d)

	// 2D array
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = 0
		}
	}
	fmt.Println("2D: ", twoD)

	// create and initialize multidimensional arrays
	twoD = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2D: ", twoD)

}
