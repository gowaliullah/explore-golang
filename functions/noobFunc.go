package main

import "fmt"

func noobFunc() {

	// in local scope --- function expression can't incoked before definde
	// aviously call after define
	// but globally working

	add(1, 1)

	add := func(a int, b int) {
		c := a + b
		fmt.Println(c)
	}

	add(2, 3)

}

func init() {
	fmt.Println("from init func")
}
