package main

func base8Char(bits byte) string {
	const base8Alphabet = "ABCDEFGH"
	if bits > byte(len(base8Alphabet)-1) {
		return ""
	}
	return string(base8Alphabet[bits])
}