package main

func crypt(plaintext, key []byte) []byte {
	cipherText := []byte{}
	for i, b := range plaintext {
		cipherText = append(cipherText, b ^ key[i])
	}
	return cipherText
}
