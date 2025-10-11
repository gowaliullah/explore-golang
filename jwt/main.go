package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	data := []byte("Hello")

	hash := sha256.Sum256(data)

	fmt.Println(hash)

	/*
		s := "098fg"

		byteArr := []byte(s)

		enc := base64.URLEncoding.WithPadding(base64.NoPadding)
		b64str := enc.EncodeToString(byteArr)

		decodedStr, err := enc.DecodeString(b64str)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(b64str)
		fmt.Println(decodedStr)

	*/

}
