package main

import "strings"

func encrypt(plaintext string, key int) string {
	return crypt(plaintext, key)
}

func decrypt(ciphertext string, key int) string {
	return crypt(ciphertext, -key)
}

func crypt(text string, key int) string {
	var result string
	for _, c := range text {
		result += getOffsetChar(c, key)
	}
	return result
}

func getOffsetChar(c rune, offset int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	charIndex := strings.IndexRune(alphabet, c)
	if charIndex == -1 {
		return ""
	}
	// Handle negative offsets correctly by adding 26 to ensure a positive result
	newIndex := (charIndex + offset) % 26
	if newIndex < 0 {
		newIndex += 26
	}
	return string(alphabet[newIndex])
}
