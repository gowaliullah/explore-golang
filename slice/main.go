package main

import "fmt"

func main() {

	// empty || nil slice

	var s []int // ptr = nil, len = 0, cap = 0

	fmt.Println("slice: ", s, "Len: ", len(s), "Cap: ", cap(s))

	// make() with len && cap
	/*
		s := make([]int, 3, 5) // [0 0 0 ] cap = 3 len = 5
		s[0] = 5
		s[2] = 4 // [5 0 4 ] cap = 3 len = 5
		s[3] = 4 //panic: runtime error: index out of range [3] with length 3
		fmt.Println("slice: ", s, "Len: ", len(s), "Cap: ", cap(s))

	*/

	// make() with len
	/*
		s := make([]int, 3) // [0 0 0 ] cap = 3 len = 3
		s[0] = 5 // [5 0 0 ] cap = 3 len = 3

		fmt.Println("slice: ", s, "Len: ", len(s), "Cap: ", cap(s))

	*/

	// second step of learning slice
	/*
		// arr := [3]int{1, 2, 3}
		// fmt.Println(arr)

		slc := []int{1, 2, 3} // slice literal -> S.L. দিয়ে করলে len & cap একই হয়।
		fmt.Println("slice: ", slc, "Len: ", len(slc), "Cap: ", cap(slc))

	*/

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
