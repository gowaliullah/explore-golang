package main

import (
	"fmt"
	"os"
)

type user struct {
	Name  string
	Age   int
	Money float64
}

type People interface {
	added(m float64) float64
}

func (obj user) added(m float64) float64 {
	obj.Money = obj.Money + m
	return obj.Money
}

func (obj user) withdrawMoney(m float64) float64 {
	obj.Money = obj.Money - m
	return obj.Money
}

type Owner interface {
	withdrawMoney(m float64) float64
}

func main() {

	var usr1 People
	usr1 = user{
		Name:  "Habi",
		Age:   19,
		Money: 1234953.434,
	}

	usr1.added(1000000)

	var usr2 Owner
	usr2 = user{
		Name:  "Habi",
		Age:   19,
		Money: 1234953.434,
	}

	usr2.withdrawMoney(8998878)

	obj, ok := usr2.(user)

	if !ok {
		fmt.Println("sorry..! usr2 is not type of user struct")
		os.Exit(1)
	}

	obj.added(12213232423445.5345)

	obj.withdrawMoney(24343.4324)

}
