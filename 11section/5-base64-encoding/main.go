package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	data := "Welcome to wonderful world of Go!"

	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encoded)

	encodedString := "V2VsY29tZSB0byB3b25kZXJmdWwgd29ybGQgb2YgR28h"
	decoded, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		log.Fatal(err)
	}

	if string(decoded) != data {
		log.Fatalf("decoded string does not match encoded data")
	}

	rawData := []byte{0xDE, 0xAD, 0xEF, 0xCA, 0xFE}
	binaryCodedToString := base64.StdEncoding.EncodeToString(rawData)

	fmt.Println(string(decoded))
	fmt.Println(string(binaryCodedToString))

	b64Str := "3q3vyv4="
	decodedB64, err := base64.StdEncoding.DecodeString(b64Str)
	if err != nil {
		log.Fatal(err)
	}

	if string(decodedB64) != data {
		log.Fatalf("decoded string does not match encoded data")
	}

	fmt.Println(string(decodedB64))
}
