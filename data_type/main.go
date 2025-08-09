package main

import "fmt"

func main() {

	var p int8 = -128
	var q uint8 = 255 // unsigned ( 0 and only positive number)
	var z float32 = 20.40
	var ha bool = false // 8 bits - 1 byte

	// byte -> alias for unit8 -> 8 bits -> 1 byte

	// rune -> alias for init32 (unicode point) -> 32 bits -> 4 bytes %c

	fmt.Println(p, q, z, ha)
}
