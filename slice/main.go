package main

import "fmt"

func main() {

	// second step of learning slice

	// arr := [3]int{1, 2, 3}
	// fmt.Println(arr)

	slc := []int{1, 2, 3} // slice literal -> S.L. দিয়ে করলে len & cap একই হয়।
	fmt.Println("slice: ", slc, "Len: ", len(slc), "Cap: ", cap(slc))

	// first step of learing slice
	/*
		arr := [6]string{"This", "is", "a", "GO", "interview", "question"} // tmi

		s := arr[1:4] // "is", "a", "GO"  // haat

		s1 := s[1:2] // "a" // ring

		fmt.Println(s1)
		fmt.Println("len: ", len(s1))
		fmt.Println("Cap: ", cap(s1))

	*/
}
