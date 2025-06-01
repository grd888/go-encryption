package main

import (
	"math/big"
)

func encrypt(m, e, n *big.Int) *big.Int {
	// Calculate m^e mod n using the Exp method
	result := new(big.Int).Exp(m, e, n)
	return result
}