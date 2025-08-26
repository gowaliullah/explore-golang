package main

import (
	"fmt"
)

func forLoop() {
	nums := [5]int{1, 2, 3, 4, 5}

	fmt.Println("For Loop")
	for i := 0; i < len(nums); i++ {
		fmt.Println(i, nums[i])
	}
}

func ForLoopWithRange() {
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Println("For Loop with range")
	for index, value := range nums {
		fmt.Println(index, value)
	}
}

func multidimensionalArray() {
	matrix := [2][2]int{
		{1, 2},
		{4, 5},
	}
	fmt.Println(matrix)
}

func arrayOfArrays() {
	arr := [][]int{
		{1, 2},
		{3, 4},
		{5},
		{6, 7, 8},
	}
	fmt.Println(arr)
}

func modifyArr(arr *[6]int) {
	arr[0] = 999
}

func main() {
	nums := [6]int{1, 2, 3, 4, 5, 6}
	modifyArr(&nums)
	fmt.Println(nums)
}
