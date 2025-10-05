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

	// as like pagination

	page := []string{"1", "2", "3", "4", "5"}
	limit := []string{"10", "20", "30"}

	params := map[string][]string{
		"page":  page,
		"limit": limit,
	}
	fmt.Println(params)

}
