// use `slices` package provides sorting for builtins and user-defined types
// work for any ordered type defined here: https://pkg.go.dev/cmp#Ordered

package main

import (
	"cmp"
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

	// custom sort

	// - by length of strings
	fruits := []string{"peach", "banana", "kiwi"}
	// custom sort function
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	// - by age of person
	type Person struct {
		name string
		age  int
	}
	people := []Person{
		{name: "Jay", age: 37},
		{name: "Trish", age: 25},
		{name: "Alex", age: 72},
	}
	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(people)
}
