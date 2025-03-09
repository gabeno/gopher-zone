package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type response1 struct {
	Page   int
	Fruits []string
}

// only exported fields will be encoded/decoded in json
// fields nust start with capital letters to be exported
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	//
	// encoding
	//

	// basic data types
	boolB, _ := json.Marshal(true)
	fmt.Println("bool:", string(boolB))

	intB, _ := json.Marshal(2)
	fmt.Println("int:", string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println("float:", string(fltB))

	strB, _ := json.Marshal("name")
	fmt.Println("string:", string(strB))

	slcB, _ := json.Marshal([]string{"a", "b", "c"})
	fmt.Println("slice:", string(slcB))

	mapB, _ := json.Marshal(map[string]int{"apples": 3, "pears": 4})
	fmt.Println("map:", string(mapB))

	// encode cutom data types
	// only include exported fields in the encoded output and by default uses
	// the names as JSON keys
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apples", "peach", "pears"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println("resp1:", string(res1B))

	// use tags on struct field to customize the encoded JSON key names
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apples", "peach", "pears"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println("resp1:", string(res2B))

	//
	// decoding
	//

	byt := []byte(`{"num":6.124,"strs":["a","b"]}`)

	// provide a variable where the json package can put the decoded data
	// `map[string]any` will hold a map os strings with abitrary data types
	var dat map[string]any

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// to use the values in the decoded map, we need to convert them into their
	// appropriate type
	num := dat["num"].(float64)
	fmt.Println(num)

	// accessing nested data requires a series of conversions
	strs := dat["strs"].([]any)
	str1 := strs[0].(string)
	fmt.Println(str1)

	// decode json into custom data types
	// ensures type-safety and eliminating the need for type assertions when
	// accessing the decoded data
	str := `{"page":1,"fruits":["apples","peach","pears"]}`
	res := response2{}
	if err := json.Unmarshal([]byte(str), &res); err != nil {
		panic(err)
	}
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// instead of bytes and strings as intermediaries between data and json in stdout,
	// we can stream JSON encoding directly to `os.Writers` like `os.stdout` or HTTP
	// response bodies
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	// streaming reads from `os.Readers` like `os.Stdin` or HTTP request bodies is done
	// with `json.Decoder`
	dec := json.NewDecoder(strings.NewReader(str))
	res1 := response2{}
	dec.Decode(&res1)
	fmt.Println(res1)
	fmt.Println(res1.Fruits[0])
}

// @TODO: read -> https://go.dev/blog/json and https://pkg.go.dev/encoding/json
