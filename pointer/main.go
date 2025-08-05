package main

// hdd -> second memory

import "fmt"

func print(nums [3]int) {
	fmt.Println(nums)
}

func print2(nums *[3]int) {
	fmt.Println(nums)
}

type User struct {
	Name     string
	Age      int
	Salary   float64
	FavFoods []string
}

func main() {

	// pointer // address // memory ( ram / hard disk )
	/*
		x := 10

		p := &x // ampersane --> address of

		*p = 1000

		fmt.Println("X:", x)
		fmt.Println("Address:", p)               // address of x
		fmt.Println("value at the address:", *p) // *  value of address
	*/
	// pass by value
	// pass by reference

	habib := User{
		Name:   "Habib",
		Age:    27,
		Salary: 0,
	}

	fmt.Println(habib)
	q := &habib

	q.Name = "Shamin"

	fmt.Println(q)

	// arr := [3]int{1, 2, 3}

	// print(arr)
	// print2(&arr)

}
