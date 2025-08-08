package main

import "fmt"

func add(x int, y int) int {
	var res int

	res = x + y
	return res
}

func main() {
	var a int = 10

	var sum = add(a, 4)

	fmt.Println(sum)

	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// defer fmt.Println(4)
	// defer fmt.Println(1)

	// fmt.Println(100)

	func() { // annonymous func
		fmt.Println("habib")
	}() // IFFE

	add := func() { // function expression
		fmt.Println("habib")
	}

	add()

}

/*


Pointing Register (Program Counter) প্রথমে প্রথম instruction address কে point করে।
CPU সেই instruction -> fetch -> decode -> execute করে।
Instruction শেষ হলে, Pointing register এক ধাপ বাড়ে, অর্থাৎ পরবর্তী instruction address এ চলে যায়।



*/
