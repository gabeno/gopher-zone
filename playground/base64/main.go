// base64 encoding/decoding

package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"

	// Go supports both standard and url-compatible base64
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))

	// using a url-compatible base64 format
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)

	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
