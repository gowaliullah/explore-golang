package main

import (
	"fmt"
	"time"
)

func printHello(num int) {
	time.Sleep(5 * time.Second)
	fmt.Println("Helllo World", num)
}

var p, q int = 10, 20

func main() {
	go printHello(1)
	go printHello(2)
	go printHello(3)
	go printHello(4)

	fmt.Println(p, " ", q)

	time.Sleep(5 * time.Second)

	fmt.Println("ha....")

}
