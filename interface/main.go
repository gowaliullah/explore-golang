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

type BankUser interface {
	WithdrawMoney(amount float64) float64
}

// receiver func or methods or behaviour
func (obj user) PrintDetails() {
	fmt.Println("Name: ", obj.Name)
	fmt.Println("Age: ", obj.Age)
	fmt.Println("Money: ", obj.Money)
}

func (obj user) ReceivedMoney(amount float64) float64 {
	obj.Money = obj.Money + amount
	return obj.Money
}

func (obj user) WithdrawMoney(amount float64) float64 {
	obj.Money = obj.Money - amount
	return obj.Money
}

func main() {
	var usr1 People
	usr1 = user{
		Name:  "Hab",
		Age:   30,
		Money: 344.423,
	}
	usr1.PrintDetails()
	usr1.ReceivedMoney(100)
}
