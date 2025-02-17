package main

import (
	"fmt"
	"time"
)

func main() {

	// basic switch statement
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("null")
	}
	fmt.Println("")

	// use commas to separate multiple expressions
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its the weekend")
	default:
		fmt.Println("its weekday")
	}
	fmt.Println("")

	// switch without expression is alternate way to express if/else logic
	// expressions can be non-constants
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("its before noon")
	default:
		fmt.Println("its after noon")
	}
	fmt.Println("")

	// type switch compares types instead of values
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an int")
		case string:
			fmt.Println("I am a string")
		default:
			fmt.Printf("Dont know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(8)
	whatAmI("true")
	whatAmI(nil)

}
