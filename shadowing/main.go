package main

import "fmt"

func main() {

	a := 21

	if true {
		a = 12
		fmt.Println("Inner the scope:", a)
	}

	fmt.Println("Outter the scope: ", a)

}
