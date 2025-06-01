package main

import (
	"math/big"
)

// Get the private exponent
func getD(e, tot *big.Int) *big.Int {
	// Calculate the modular multiplicative inverse of e (mod tot)
	d := new(big.Int).ModInverse(e, tot)
	return d
}
