package main

import "fmt"

func main() {
	// create an empty map: map[key-type]value-type
	m := make(map[string]int)

	// set key-value pairs
	m["k1"] = 1
	m["k2"] = 2
	fmt.Println("map:", m)

	// get a value
	fmt.Println("get:", m["k1"])

	// if key does not exist, zero value of the type is returned
	fmt.Println("get:", m["k3"])

	// len returns the number of key-value pairs
	fmt.Println("len:", len(m))

	// delete removes key-value pairs
	delete(m, "k2")
	fmt.Println("del:", m)

	// remove all ky-value pairs
	clear(m)
	fmt.Println("clear:", m)

	// is value present in map, check with second return value
	// useful in disambiguating between missing keys and keys with zero values
	val, present := m["k3"]
	fmt.Println("val:", val, "present:", present)

	// declare and initialise
	n := map[string]int{"a": 1, "b": 2}
	fmt.Println("decl:", n)

	// @todo: check maps package -> https://pkg.go.dev/maps
}
