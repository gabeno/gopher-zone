package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// match with a string pattern
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// alternatively, compile an optimized `regexp` struct
	r, _ := regexp.Compile("p([a-z]+)ch")

	// many methods are available on these structs
	fmt.Println("MatchString:", r.MatchString("peach"))
	fmt.Println("FindString:", r.FindString("peach punch"))
	fmt.Println("FindStringIndex:", r.FindStringIndex("peach punch"))
	fmt.Println("FindStringSubmatch:", r.FindStringSubmatch("peach punch"))
	fmt.Println("FindStringSubmatchIndex:", r.FindStringSubmatchIndex("peach punch"))
	fmt.Println("FindAllString:", r.FindAllString("peach punch pinch", -1))
	fmt.Println("FindAllString:", r.FindAllString("peach punch pinch", 1)) // liimit number of matches

	// with []byte arguments - drop "String" from function name
	fmt.Println("Match:", r.Match([]byte("peach")))

	// use `MustCompile` when creating global variables with regular expressions
	// it panics instead of returning an error
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	fmt.Println("ReplaceAllString:", r.ReplaceAllString("a peach", "<fruit>")) // mask PII :)
	// transofmr output with `Func` variant
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}

// @TODO: check more methods in docs: https://pkg.go.dev/regexp
