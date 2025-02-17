package main

import "fmt"

func main() {
	// uninitialised slice has length zero and equals to nil
	var s []string
	fmt.Println("init slice: ", s, len(s) == 0, s == nil)

	// empty slice with non-zero length
	// by default capacity == length
	// make takes an arg to determine capacity
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// set and get values
	s[0] = "a"
	s[1] = "b"
	fmt.Println("set:", s)
	fmt.Println("get@1:", s[0])
	fmt.Println("get@2:", s[1])
	fmt.Println("get@3:", s[2])

	// length of slice
	fmt.Println("len:", len(s))

	// append to slice
	// return a slice with one or more values
	s = append(s, "c")
	s = append(s, "d", "e", "f")
	fmt.Println("append:", s)

	// copy slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	// slice operator: s[low:high]
	fmt.Println("sliced", s[2:5])       // from low to high, excluding high
	fmt.Println("slice up to:", s[:5])  // up to high, excluding high
	fmt.Println("slice up from", s[2:]) // up from and including

	// declare and initialize a slice
	t := []string{"a", "b", "c"}
	fmt.Println("decl:", t)

	// two d slice
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D:", twoD)

	// @todo: check slices package -> https://pkg.go.dev/slices
	// @todo: read slices internals -> https://go.dev/blog/slices-intro
}
