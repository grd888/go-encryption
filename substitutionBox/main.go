package main

import "errors"

func sBox(b byte) (byte, error) {
	// Only use the last 4 bits of the input
	input := b & 0x0F // Mask to get only the last 4 bits

	// Lookup table based on the test cases and the pattern in the assignment
	// The index represents the 4-bit input, the value is the 2-bit output
	lookupTable := []byte{
		0b00, // 0000 -> 00
		0b10, // 0001 -> 10
		0b01, // 0010 -> 01
		0b11, // 0011 -> 11
		0b10, // 0100 -> 10
		0b01, // 0101 -> 01
		0b11, // 0110 -> 11
		0b01, // 0111 -> 01
		0b01, // 1000 -> 01
		0b11, // 1001 -> 11
		0b00, // 1010 -> 00
		0b10, // 1011 -> 10
		0b11, // 1100 -> 11
		0b01, // 1101 -> 01
		0b10, // 1110 -> 10
		0b00, // 1111 -> 00
	}

	// Check if input is within valid range (0-15)
	if int(input) >= len(lookupTable) {
		return 0, errors.New("invalid input")
	}

	return lookupTable[input], nil
}
