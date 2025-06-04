package main

import (
	"crypto/sha256"
	"fmt"
)

func checksumMatches(message string, checksum string) bool {
	hasher := sha256.New()
	hasher.Write([]byte(message))
	return fmt.Sprintf("%x", hasher.Sum(nil)) == checksum
}
