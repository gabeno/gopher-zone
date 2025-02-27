// Strings - a read-only slice of bytes encoded in utf8
// Rune - concept of a character, an integer that represents a unicode code point

// @todo: read about strings, runes and characters -> https://go.dev/blog/strings

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// hello in thai
	// go string literals are utf8 encoded
	const s = "สวัสดี"

	// raw string literal vs string literal
	fmt.Println(`
<html>
	<body>"Hello"</body>
</html>`) // raw string literal: no need for escape sequences, more readable in code
	fmt.Println("<html>\n\t<body>\"Hello\"</body>\n</html>") // string literal: need escape sequences,

	// strings are equivalent to []byte
	// length of raw bytes stored in s
	// len returns number of bytes in a string
	fmt.Println("len:", len(s))

	// raw byte values at each index
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i]) // hex values for bytes
	}
	fmt.Println()

	// count runes in string aka code points
	// runtime depends on size of string coz decodes each utf8 rune sequentially
	// characters are represented by utf8 code points that span multiple bytes
	fmt.Println("rune count:", utf8.RuneCountInString(s))

	// rune value with its offset
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
	fmt.Println()

	// rune value with offset using DecodeRuneInString
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// passing a rune value to a function
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	// values enclosed in single quotes are rune literals
	// compare rune value to rune literal
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
