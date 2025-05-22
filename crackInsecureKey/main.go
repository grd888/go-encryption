package main

import (
	"errors"
	"math"
	"strings"
)

func findKey(encrypted []byte, decrypted string) ([]byte, error) {
	for i := 0; i < int(math.Pow(2, 24)); i++ {
    key := intToBytes(i)
		ciphertext := crypt(encrypted, key)
		// cast ciphertext to string
		ciphertextStr := string(ciphertext)
		// check if the ciphertext equals the decrypted string
		if strings.Contains(ciphertextStr, decrypted) {
			return key, nil
		}	
	}	
	return nil, errors.New("Key not found")
}	