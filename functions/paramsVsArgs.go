package main

import "fmt"

func add(a, b int) { // parametar a & b. when received called parameter

	fmt.Println(1 + 2)

}

func paramsVsArgs() {
	add(1, 2) // arguments 1 2 when passed called agguments
}

func init() {
	fmt.Println("from init func")
}
