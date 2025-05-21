package main

import (
	"fmt"
	"strings"
)

func getHexString(b []byte) string {
	var result []string
	for _, v := range b {
		result = append(result, fmt.Sprintf("%02x", v))
	}
	return strings.Join(result, ":")
}

func getBinaryString(b []byte) string {
	var result []string
	for _, v := range b {
		result = append(result, fmt.Sprintf("%08b", v))
	}
	return strings.Join(result, ":")
}
