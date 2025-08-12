package main

import "fmt"

func sum(a, b int) (result int) {
	result = a + b

	return //  default return result
}

func main() {

	res := sum(3, 4)

	fmt.Println(res)

}
