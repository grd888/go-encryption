package main

import (
	"crypto/sha256"
	"encoding/hex"
)

func hmac(message, key string) string {
	byteKey := []byte(key)
	length := len(byteKey)
	midPoint := length / 2
	
	// Split the key into two halves
	// Second half should be the larger half if key's length is odd
	keyFirstHalf := byteKey[:midPoint]
	keySecondHalf := byteKey[midPoint:]
	
	// First hash: sha256(keySecondHalf + message)
	innerHash := sha256.New()
	innerHash.Write(keySecondHalf)
	innerHash.Write([]byte(message))
	innerDigest := innerHash.Sum(nil)
	
	// Final hash: sha256(keyFirstHalf + innerDigest)
	outerHash := sha256.New()
	outerHash.Write(keyFirstHalf)
	outerHash.Write(innerDigest)
	
	// Return the final hash as a lowercase hex string
	return hex.EncodeToString(outerHash.Sum(nil))
}
