// compute sha256 hashes in Go
// sha256 hashes are used to compute short identities for binary or text blobs.

package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"

	h := sha256.New()

	// `Write` expects bytes
	h.Write([]byte(s))

	// hash result as a byte slice
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}

// @TODO: for cryptogaphically secure hashes -> https://en.wikipedia.org/wiki/Cryptographic_hash_function
