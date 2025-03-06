// use `slices` package provides sorting for builtins and user-defined types
// work for any ordered type defined here: https://pkg.go.dev/cmp#Ordered

package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []string{"b", "a", "c"}
	slices.Sort(s)
	fmt.Println("strings:", s)

	n := []int{7, 9, 4, 2}
	slices.Sort(n)
	fmt.Println("ints", n)

	n_isSorted := slices.IsSorted(n)
	fmt.Println("sorted n", n_isSorted)
}
