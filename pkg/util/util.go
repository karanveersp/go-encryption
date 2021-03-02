package util

import (
	"fmt"
)

// KeyToBytes takes a string and returns a 32
// bit padded byte array. If the string
// has 32 characters then the byte array is not
// padded.
func KeyToBytes(key string) ([]byte, error) {
	if len(key) > 32 {
		return nil, fmt.Errorf("key should be <= 32 characters")
	}
	barr := [32]byte{}
	for i, v := range []byte(key) {
		barr[i] = v
	}
	return barr[:], nil
}
