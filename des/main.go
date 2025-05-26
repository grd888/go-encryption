package main

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"io"
)

func encrypt(key, plaintext []byte) ([]byte, error) {
	// Create a new DES cipher block
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Pad the plaintext to be a multiple of the block size
	paddedPlaintext := padMsg(plaintext, des.BlockSize)

	// Create a buffer to hold the IV and ciphertext
	ciphertext := make([]byte, des.BlockSize+len(paddedPlaintext))

	// Generate a random IV and put it at the beginning of the ciphertext
	iv := ciphertext[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// Create the CBC encrypter and encrypt the data
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[des.BlockSize:], paddedPlaintext)

	return ciphertext, nil
}

func padMsg(plaintext []byte, blockSize int) []byte {
	// If the plaintext is already a multiple of blockSize, we still need to add a full block
	// to indicate that padding was added (this is PKCS#7 style padding but with zeros)
	paddingNeeded := blockSize - (len(plaintext) % blockSize)
	if paddingNeeded == 0 {
		paddingNeeded = blockSize
	}

	// Create a new slice with enough space for the padding
	padded := make([]byte, len(plaintext)+paddingNeeded)
	copy(padded, plaintext)
	
	// Pad with zeros using the helper function
	return padWithZeros(padded[:len(plaintext)], len(padded))
}
