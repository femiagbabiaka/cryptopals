package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	// Our input string in bytes.
	hexByte := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	hexToBase64(hexByte)
}

func hexToBase64(b []byte) []byte {

	// Allocate space for our decoded bytes to our var.
	decodedBytes := make([]byte, hex.DecodedLen(len(b)))

	_, err := hex.Decode(decodedBytes, b)

	if err != nil {
		panic(err)
	}

	// Same.
	encodedBytes := make([]byte, base64.StdEncoding.EncodedLen(len(decodedBytes)))
	// Encode our bytes to base 64.
	base64.StdEncoding.Encode(encodedBytes, decodedBytes)

	fmt.Println(string(b))
	fmt.Println(string(encodedBytes))

	return b

}
