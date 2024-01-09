/*
golang challenge:

write a function walk(x interface{}, fn func(string)) which takes a struct x
and calls fn for all strings fields found inside.

difficulty level: recursively.

*/

package walk

func walk(x interface{}, fn func(input string)) {
	fn("hey yoda")
}
