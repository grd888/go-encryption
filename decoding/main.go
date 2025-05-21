package main

import (
	"encoding/hex"
	"strings"
)

func getHexBytes(s string) ([]byte, error) {
	hexStrings := strings.Split(s, ":")
	var result []byte
	for _, v := range hexStrings {
		hexByte, err := hex.DecodeString(v)
		if err != nil {
			return nil, err
		}
		result = append(result, hexByte...)
	}
	return result, nil
}