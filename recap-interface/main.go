package main

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

}
