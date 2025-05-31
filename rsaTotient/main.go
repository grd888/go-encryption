package main

import (
	"math/big"
)

func getTot(p, q *big.Int) *big.Int {
	// Create copies of p and q to avoid modifying the originals
	p1 := new(big.Int).Sub(p, big.NewInt(1))
	q1 := new(big.Int).Sub(q, big.NewInt(1))
	
	// Calculate (p-1)*(q-1)
	t := new(big.Int).Mul(p1, q1)
	return t
}

func getE(tot *big.Int) *big.Int {
	// Create a big.Int for the lower bound (2)
	two := big.NewInt(2)
	
	// Create a big.Int for the upper bound (tot-1)
	upper := new(big.Int).Sub(tot, big.NewInt(1))
	
	// Create a range for the random number: [0, tot-2)
	range1 := new(big.Int).Sub(upper, two)
	
	// Loop until we find a suitable e
	for {
		// Generate a random number in range [0, tot-2)
		random := new(big.Int)
		random.Rand(randReader, range1)
		
		// Add 2 to get a number in range [2, tot-1)
		e := new(big.Int).Add(random, two)
		
		// Check if e and tot are coprime (gcd = 1)
		if gcd(e, tot).Cmp(big.NewInt(1)) == 0 {
			return e
		}
	}
}
