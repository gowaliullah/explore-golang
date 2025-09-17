package main

import "fmt"

type user struct {
	Name  string
	Age   int
	Money float64
}

// only signatures of the functions
type People interface { // pure abstructs
	PrintDetails()
	ReceivedMoney(amount float64) float64
}

// receiver func or methods or behaviour
func (obj user) PrintDetails() {
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
	usr1.PrintDetails()
}
