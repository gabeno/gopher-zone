// embedding of structs and interfaces to express a more seamless composition of types

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// container embeds a base
type container struct {
	base
	str string
}

func main() {
	// initialize structs explicitly when creating structs with literals
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// access base fields directly on co
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// spell out the full path
	fmt.Printf("co={num: %v, str: %v}\n", co.base.num, co.str)

	// container embeds base so methods of base become methods of container
	fmt.Println("describe:", co.describe())

	// embedding structs with methods can be used to bestow interface implementations on other structs
	// container now implements describer interface because it embeds base
	type describer interface {
		describe() string
	}
	var d describer = co
	fmt.Println("describer:", d.describe())
}
