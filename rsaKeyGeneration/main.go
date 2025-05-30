package main

import (
	"math/big"
)

func generatePrivateNums(keysize int) (*big.Int, *big.Int) {
	p, _ := getBigPrime(keysize)
	q, _ := getBigPrime(keysize)
	return p, q
}

// Calculate n = p * q
func getN(p, q *big.Int) *big.Int {
	z := new(big.Int)
	z.Mul(p, q)
	return z
}
