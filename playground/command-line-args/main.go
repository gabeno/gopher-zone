package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

// go run command-line-args/main.go a b c d
// [/var/folders/fb/s5rmv2xs31988y78v074x63r0000gn/T/go-build2860757870/b001/exe/main a b c d]
// [a b c d]
// c
