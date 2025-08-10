package main

import "fmt"

func a() {

	i := 0

	fmt.Println("First: ", i)

	defer fmt.Println("Second: ", i)

	i++

	fmt.Println("Third: ", i)
	defer fmt.Println("Forth: ", i)
}

func main() {
	a()
}
