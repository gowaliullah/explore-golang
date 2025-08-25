package main

import "fmt"

func main() {

	a := 21
	b := 21

	if true {
		a := 12                     // new variable -> shadowing
		b = 12                      // assign -> so not shadowing
		fmt.Println("Inner:", a, b) // 12, 12
	}
	fmt.Println("Outter: ", a, b) // 21, 12
}
