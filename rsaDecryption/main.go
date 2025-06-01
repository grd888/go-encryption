package main

import (
	"math/big"
)

func decrypt(c, d, n *big.Int) *big.Int {
	// Calculate c^d mod n using the Exp method
	result := new(big.Int).Exp(c, d, n)
	return result
}
