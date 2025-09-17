package main

import "fmt"

type user struct {
	Name  string
	Age   int
	Money float64
}

type User interface {
}

func PrintDetails(obj user) {
	fmt.Println("Name: ", obj.Name)
	fmt.Println("Age: ", obj.Age)
	fmt.Println("Money: ", obj.Money)
}

func main() {
	usr1 := user{
		Name:  "Hab",
		Age:   30,
		Money: 344.423,
	}
	PrintDetails(usr1)
}
