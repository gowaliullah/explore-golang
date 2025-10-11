package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "098fg"

	byteArr := []byte(s)

	enc := base64.URLEncoding.WithPadding(base64.NoPadding)
	b64str := enc.EncodeToString(byteArr)

	decodedStr, err := enc.DecodeString(b64str)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(b64str)
	fmt.Println(decodedStr)

}
