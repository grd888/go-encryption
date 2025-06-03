package main

import (
	"crypto/sha256"
	"fmt"
	"hash"
)

type hasher struct {
	hash hash.Hash
}

// ?
func newHasher() *hasher {
	h := sha256.New()
	return &hasher{hash: h}
}

func (h *hasher) Write(password string) {
	h.hash.Write([]byte(password))
}

func (h *hasher) GetHex() string {
	return fmt.Sprintf("%x", h.hash.Sum(nil))
}