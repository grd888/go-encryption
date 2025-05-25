package main

import (
	"crypto/aes"
	"crypto/des"
	"errors"
)

func getBlockSize(keyLen, cipherType int) (int, error) {
	switch cipherType {
	case typeAES:
		cipher, err := aes.NewCipher(make([]byte, keyLen))
		if err != nil {
			return 0, err
		}
		return cipher.BlockSize(), nil
	case typeDES:
		cipher, err := des.NewCipher(make([]byte, keyLen))
		if err != nil {
			return 0, err
		}
		return cipher.BlockSize(), nil
	default:
		return 0, errors.New("invalid cipher type")
	}
}
