package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 64 is how many bits of precision to parse
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// convenience function for basic base-10 `int` parsing
	k, _ := strconv.Atoi("123")
	fmt.Println(k)

	// error out when input is bad
	_, e := strconv.Atoi("what")
	fmt.Println(e)
}
