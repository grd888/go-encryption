package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

func genKeys() (pubKey *ecdsa.PublicKey, privKey *ecdsa.PrivateKey, err error) {
	privKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	pubKey = &privKey.PublicKey
	return pubKey, privKey, nil
}
