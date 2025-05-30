package main

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/sha256"
)

func encrypt(pubKey *rsa.PublicKey, msg []byte) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, msg, nil)
}
