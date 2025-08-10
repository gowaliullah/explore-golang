package main

import "fmt"

// A callback function that checks if the number is odd or even
func machine(number int) {

	if number < 0 {
		fmt.Println(number, "digit is a INVALID number")
	} else if number%2 == 1 {
		fmt.Println(number, "digit is a ODD number")
	} else {
		fmt.Println(number, "digit is a ENVEN number")
	}

}

// A higher-order function that takes a function as an argument
func sum(num int) {
	machine(num)
}

func main() {

	// Prompting user to input a number
	var takeNum int

	fmt.Print("Enter a number: ")
	fmt.Scan(&takeNum) // Correctly taking user input and storing it in takeNum

	// Calling sum with the user's input
	sum(takeNum)
}
