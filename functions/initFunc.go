package main

import "fmt"

var a = 10

// var b = 20

// func main(){
// 	var sum = a + b
// }

// func main(){
// 	fmt.Println("Hello init func")
// }

// func init(){
// 	fmt.Println("This is the function that is executed first")
// }

func initFunc() {
	fmt.Println(a)
}

func init() {
	// var a = 5
	fmt.Println("from init func", a)
	a = 20
}
