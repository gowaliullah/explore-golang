package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "098fg"

	byteArr := []byte(s)

	enc := base64.URLEncoding.WithPadding(base64.NoPadding)
	str := enc.EncodeToString(byteArr)

	fmt.Println(str)

}
