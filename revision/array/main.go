package main

import "fmt"

func main() {
	nums := [5]int{1, 2, 3, 4, 5}

	fmt.Println("For Loop")
	for i := 0; i < len(nums); i++ {
		fmt.Println(i, nums[i])
	}

	fmt.Println("For Loop with range")
	for index, value := range nums {
		fmt.Println(index, value)
	}

}
