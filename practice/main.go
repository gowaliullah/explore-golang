package main

import "fmt"

func main() {
	a := []int{1, 2, 3} // len = 3, cap = 3
	a = append(a, 4)    // len = 4, cap = 6
	a = append(a, 5)    // len = 5, cap = 6

	b := a[3:]       // len = 2, cap = 3
	b = append(b, 6) // len = 3, cap = 3
	a[3] = 77
	b = append(b, 7) // len = 4, cap = 6
	b = append(b, 8) // len = 5, cap = 6

	a = append(a, 9) // len = 6, cap = 6

	fmt.Println(b)
	fmt.Println(a)

	// fmt.Println(len(a), cap(a)) // len = 6, cap = 6
	// fmt.Println(len(b), cap(b)) // len = 5, cap = 8
}
