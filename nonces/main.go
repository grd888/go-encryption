package main
import "math"

// nonceStrength returns the number of bits of entropy in the nonce.
func nonceStrength(nonce []byte) int {
	nonceLength := float64(len(nonce) * 8)
	return int(math.Pow(2, nonceLength))
}
