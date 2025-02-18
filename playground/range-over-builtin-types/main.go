package main

import "fmt"

func main() {
	// slice
	nums := []int{1, 2, 3, 4}
	for idx, num := range nums {
		fmt.Println("idx:", idx, "num:", num)
	}

	// map
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m { // key/value
		fmt.Printf("%s -> %d\n", k, v)
	}
	for k := range m { // keys only
		fmt.Println("k:", k)
	}

	// string
	// index - starting byte index of the rune
	// rune - the rune itself
	for index, rune := range "go" {
		fmt.Println(index, rune)
	}
}
