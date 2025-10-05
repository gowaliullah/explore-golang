package main

import "fmt"

func main() {

	s := map[int]int{
		1: 1,
		2: 10,
		3: 5,
	}

	val := s[2]
	fmt.Println(val)

	s1 := map[string]int{
		"page":  1,
		"limit": 10,
	}

	val2 := s1["page"]

	fmt.Println(val2)
}
