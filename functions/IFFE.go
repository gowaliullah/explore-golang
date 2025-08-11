package main

import "fmt"

// standard (named) function
func add(a int, b int) {
	fmt.Println(a + b)
}

func IFFE() {
	// We invoked the add function here
	add(4, 5)

	// anonymous function
	// immediately invoked function

	func(a int, b int) {
		var sum = a + b
		fmt.Println(sum)
	}(1, 2)

	// expression
	x := 50

	// if expression / if block
	if x > 10 {
		fmt.Println("x is greater than 10")
	}
}

func init() {
	fmt.Println("I'll be called first")
}
