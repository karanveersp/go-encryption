package util

import (
	"fmt"
)

// PadZeroes32 takes a string and returns a 32
// bit padded array. If the string
// has 32 characters then no padding occurs.
func PadZeroes32(data string) ([]byte, error) {
	if len(data) > 32 {
		return nil, fmt.Errorf("len data should be <= 32 characters: len %v", len(data))
	}
	barr := [32]byte{}
	for i, v := range []byte(data) {
		barr[i] = v
	}
	return barr[:], nil
}
