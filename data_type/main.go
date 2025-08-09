package main

import "fmt"

func main() {

	var p int8 = -128
	var q uint8 = 255 // unsigned ( 0 and only positive number)
	var z float32 = 20.40323
	var ha bool = false // 8 bits - 1 byte

	haa := '💖'

	var hab rune = '💖'

	fmt.Println(p, q, z, ha, haa)
	fmt.Printf("%d\n", p)
	fmt.Printf("%f\n", z)
	fmt.Printf("%2f\n", z)
	fmt.Printf("%c\n", hab)
	fmt.Printf("%v\n", ha)
	fmt.Printf("%T\n", hab)

	// byte -> alias for unit8 -> 8 bits -> 1 byte

	// rune -> alias for init32 (unicode point) -> 32 bits -> 4 bytes -> rune ar format holo: %c (without english any charactar is rune/রুন)

}
