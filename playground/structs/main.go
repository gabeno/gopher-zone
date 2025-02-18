package main

import "fmt"

type person struct {
	name string
	age  int8
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// create a new struct
	fmt.Println(person{"Bob", 20})

	// use name fields when initializing a struct
	fmt.Println(person{name: "Tish", age: 24})

	// ommited fields are zero valued
	fmt.Println(person{name: "gabe"})
	fmt.Println(person{age: 35})

	// an & yields a pointer to the struct
	fmt.Println(&person{"pete", 22})

	// idiomatic to use constructor function
	fmt.Println(newPerson("theo"))

	// access struct fields with a dot
	s := person{"will", 27}
	fmt.Println(s.age)

	// use dot with pointer
	sPtr := &s
	fmt.Println(sPtr.age)

	// structs are mutable
	s.age = 55
	fmt.Println(s.age)
	fmt.Println(sPtr.age)

	// if struct type is used for a single value, use anonymous struct type
	// important technique for table-driven tests
	dog := struct {
		name   string
		isGood bool
	}{
		"rex",
		true,
	}
	fmt.Println(dog)
}
