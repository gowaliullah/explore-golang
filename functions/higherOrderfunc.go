package main

import "fmt"

//  ## higher order function e jodi kono function K pass kora hoy, setai callBack function
//  ## je function kono function K receive kore as a parametar or retun kore or dui tai kore tahole seta heigher order function
//  ## higher order function onno name FIRST CLASS function

// ## অফিসে সবাইকে ১০ টার আগে আসতে হবে। যে জবের ইন্টারভিউ দিতে চাই তাকেও আসতে হবে

/*

any one of the following three
	1. parameter --> function
	2. function return
	3. both

*/

// func processingOperation(a int, b int, op func(c int, d int)) {
// 	op(a, b)
// }

//  ## higher order function e jodi kono function K pass kora hoy, setai callBack function
//  ## je function kono function K receive kore as a parametar or retun kore or dui tai kore tahole seta heigher order function treated as first class citizen

func call() func(x, y int) {
	return add
}

func add(x, y int) {
	fmt.Println(x + y)
}

// a := 1

func higherOrderfunc() {
	sum := call() // add assign in sum
	sum(3, 4)
}

func init() {
	fmt.Println("from init func")
}
