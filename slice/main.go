package main

import "fmt"

func changeSlice(a []int) []int {
	a[0] = 10
	a = append(a, 11)
	return a
}

func main() {

	x := []int{1, 2, 3, 4, 5}

	x = append(x, 6)
	x = append(x, 7)

	a := x[4:]

	y := changeSlice(a)

	fmt.Println(x)
	fmt.Println(y)

	// fmt.Println("slice: ", a, "Len: ", len(a), "Cap: ", cap(a)) // [10 2 3 5] Len:  4 Cap:  4

	/*
		var x []int

		x = append(x, 1) // [1], len = 1, cap = 1
		x = append(x, 2) // [1, 2], len = 2, cap = 2
		x = append(x, 3) // [1, 2, 3], len = 3, cap = 4

		y := x

		x = append(x, 4)
		y = append(y, 5)

		x[0] = 10

		fmt.Println("slice: ", x, "Len: ", len(x), "Cap: ", cap(x)) // [10 2 3 5] Len:  4 Cap:  4
		fmt.Println("slice: ", y, "Len: ", len(y), "Cap: ", cap(y)) // [10 2 3 5] Len:  4 Cap:  4

	*/

	// empty || nil slice
	/*
		var s []int // ptr = nil, len = 0, cap = 0

		s = append(s, 1)
		s = append(s, 1)
		s = append(s, 1)

		fmt.Println("slice: ", s, "Len: ", len(s), "Cap: ", cap(s))

	*/

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

/*

1. slice from existing array
2. slice from slice a array
3. slice literal
4. make function with length
5. make function with length and capacity
5. empty or nil slice
6. slice underlying array rule --> 1024 -> 100% then 25% increase


*/
