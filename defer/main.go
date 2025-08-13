// defer fucntion ---> under the hood closer

// parent & closer jodi exists kore tahole closer variable ar address refer share kore

// defer list pointer

// defer --> linked list data stracture use kore

/*

named returned values rules

1. all code executed
2. defer function store kora hobe magic box e
3. return --> all defer fucntion execute korbe
4. return korbe named variables gular values


just return types
1. all code executed
2. defer function store kora hobe magic box e
3. return values are evaluated at this time (store the return value) // ( static )
4. all defer functions execute korbe


*/

package main

import "fmt"

func calculate() (result int) {
	fmt.Println("first", result)

	sum := func() {
		result = result + 10
		fmt.Println("defer", result)
	}

	defer sum()

	result = 5
	fmt.Println("second", result)

	return

}

func cal() int {

	result := 0

	fmt.Println("first", result)

	sum := func() {
		result = result + 10
		fmt.Println("defer", result)
	}

	defer sum()

	result = 5
	fmt.Println("second", result)

	return result

}

func normalReturn() (int, int) {

	num := 0

	x := 20

	fmt.Println("প্রথম", num) // 0

	sum := func() {

		if x > 5 {
			x = 30
		}

		num = num + 10
		fmt.Println("ডিফার", num) // 15
	}

	defer sum()

	num = 5
	fmt.Println("দ্বিতীয়", num) // 5

	return num, x // 5

}

func namedReturnedValue() (result int) {

	fmt.Println("প্রথম", result) // 0

	x := 20

	sum := func() {

		if x > 5 {
			x = 30
		}

		result = result + 10
		fmt.Println("ডিফার", result) // 15
	}
	defer sum() // nil

	result = 5

	p := func(a int) {
		fmt.Println("সবার শেষে এসে, সবার আগে যাবো", a) // 5
	}

	defer p(result)

	/*
			getAvgSalary := func () int {

			result 11
		}


		defer fmt.Println(1) // 1 -> nil
		defer fmt.Println(1)
		defer fmt.Println(1)
		defer fmt.getAvgSalary(1) // result (11) -> 127
		defer fmt.Println(1)      // 1 -> 128

	*/

	fmt.Println("দ্বিতীয়", result) // 5

	return // 15,
}

func main() {
	res := namedReturnedValue()                                   // 15
	fmt.Println("আমি মেইন, আমার ছাড়া কাম হইবো নারে পাগলা..", res) // 0, 5, 5, 15, 15
}
