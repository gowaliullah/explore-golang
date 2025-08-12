package main

import "fmt"

func sum(a, b int) (result int) { // named return value
	result = a + b

	return //  default return result
}

func sum2(a, b int) int {
	result := a + b
	return result
}

func main() {

	res := sum(3, 4)

	fmt.Println(res)

}
