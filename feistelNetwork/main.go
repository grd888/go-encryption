package main

func feistel(msg []byte, roundKeys [][]byte) []byte {
	// Split the message into left and right halves
	half := len(msg) / 2
	lhs := msg[:half]
	rhs := msg[half:]

	// Process each round
	for _, key := range roundKeys {
		// Calculate next right-hand side: hash(rhs + key) XOR lhs
		nextRHS := xor(lhs, hash(rhs, key, len(rhs)))
		// Next left-hand side is current right-hand side
		lhs, rhs = rhs, nextRHS
	}

	// Concatenate right and left halves (right first, then left)
	return append(rhs, lhs...)
}
