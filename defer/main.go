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

func normalReturn() int {

	result := 0

	fmt.Println("প্রথম", result) // 0

	sum := func() {
		result = result + 10
		fmt.Println("ডিফার", result) // 15
	}

	defer sum()

	result = 5
	fmt.Println("দ্বিতীয়", result) // 5

	return result // 5

}

func namedReturnedValue() (result int) {

	fmt.Println("প্রথম", result) // 0

	sum := func() {
		result = result + 10
		fmt.Println("ডিফার", result) // 15
	}
	defer sum()

	result = 5

	p := func(a int) {
		fmt.Println("সবার শেষে এসে, সবার আগে যাবো", a) // 5
	}

	defer p(result)
	fmt.Println("দ্বিতীয়", result) // 5

	return
}

func main() {
	res := namedReturnedValue()                                   // 15
	fmt.Println("আমি মেইন, আমার ছাড়া কাম হইবো নারে পাগলা..", res) // 0, 5, 5, 15, 15
}
