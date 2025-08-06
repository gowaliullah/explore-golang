package main

import "fmt"

func main() {
	arr := [6]string{"This", "is", "a", "GO", "interview", "question"}

	s := arr[1:4]

	fmt.Println(s)
	fmt.Println("len: ", len(arr))
	fmt.Println("Cap: ", cap(arr))
}
