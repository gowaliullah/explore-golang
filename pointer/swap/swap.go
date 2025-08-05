package main

import "fmt"

func changeValue(num int) { // pass by value
	num = 100
}

func updateVaue(num *int) { // pass by reference
	*num = 100
}

func main() {
	a := 50
	changeValue(a)
	fmt.Println("a =", a) // Output: a = 50

	updateVaue(&a)
	fmt.Println("a =", a) // Output: a = 100
}
