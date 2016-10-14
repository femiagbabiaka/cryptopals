package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	encodedString := []byte("1c0111001f010100061a024b53535009181c")

	fmt.Println(string(encodedString))

	firstBytes := make([]byte, hex.DecodedLen(len(encodedString)))

	_, err := hex.Decode(firstBytes, encodedString)

	if err != nil {
		panic(err)
	}

	secondEncodedString := []byte("686974207468652062756c6c277320657965")

	secondBytes := make([]byte, hex.DecodedLen(len(secondEncodedString)))

	_, err2 := hex.Decode(secondBytes, secondEncodedString)

	if err2 != nil {
		panic(err2)
	}

	length := len(firstBytes)

	if len(secondBytes) < length {
		length = len(secondBytes)
	}

	result := make([]byte, length)

	myShittyXOR(result, firstBytes, secondBytes)

	resBytes := make([]byte, hex.EncodedLen(len(result)))

	hex.Encode(resBytes, result)

	fmt.Println(string(resBytes))
}

func myShittyXOR(res, a []byte, b []byte) {
	length := len(a)

	if len(b) < length {
		length = len(b)
	}

	for i := 0; i < length; i++ {
		res[i] = a[i] ^ b[i]
	}
}
