/*
golang challenge:

write a function walk(x interface{}, fn func(string)) which takes a struct x
and calls fn for all strings fields found inside.

difficulty level: recursively.

*/

package walk

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if field.Kind() == reflect.String {
			fn(field.String())
		}

		if field.Kind() == reflect.Struct {
			walk(field.Interface(), fn)
		}
	}
}
