package main

import (
	"fmt"
	"time"
)

func printHello(num int) {
	fmt.Println("Helllo World", num)
}

var p, q int = 10, 20

func main() {
	go printHello(1)
	go printHello(2)
	go printHello(3)
	go printHello(4)

	fmt.Println(p, " ", q)

	time.Sleep(1 * time.Millisecond)
}
