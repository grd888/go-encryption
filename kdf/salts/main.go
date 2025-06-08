package main

import (
	"crypto/rand"
	"crypto/sha256"
)

func generateSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPassword(password, salt []byte) []byte {
	data := append(password, salt...)
	hash := sha256.Sum256(data)
	return hash[:]
}
