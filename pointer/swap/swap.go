package main

import "fmt"

// func swap(x, y int) {
// 	temp := x
// 	x = y
// 	y = temp
// }

// func main() {
// 	a, b := 1, 2
// 	swap(a, b)
// 	fmt.Println(a, b) // still 1 2
// }

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b) // 2 1
}
