package main

import (
	"math/rand"
	"fmt"
)

func generateRandomKey(length int) (string, error) {
	randReader := rand.New(rand.NewSource(0))
	p := make([]byte, length)
	_, err := randReader.Read(p)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", p), nil
}
