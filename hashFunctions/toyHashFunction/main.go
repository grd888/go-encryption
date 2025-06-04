package main

import (
	"math/bits"
)
func hash(input []byte) [4]byte {
	for i := range input {
		input[i] = bits.RotateLeft8(input[i], 3)
	}
	for i := range input {
		input[i] = input[i] << 2 
	}
	final := make([]byte, 4)
	for i := range input {
		final[i%4] = final[i%4] ^ input[i]
	}
	return [4]byte(final)
}
