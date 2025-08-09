package main

//  নোট: যদি কোনো স্ট্রাক্ট ফাংশন থেকে রিটার্ন হয় বা ক্লোজারে ব্যবহার হয়, তাহলে এটি স্ট্যাকের বদলে হিপে 		স্থানান্তরিত হতে পারে।
// ১০০% — Go-তে স্ট্রাক্ট হলো ব্যবহারকারীর তৈরি একটি ডেটাটাইপ।

// struct code segment e thake...
// struct can be allocated on either the stack or the heap, depending on a process known as escape analysis performed by the Go compiler.

import "fmt"

type User struct {
	Name string // member variable or property
	Age  int
}

func printUserDetails(usr User) {
	fmt.Println(usr.Name)
	fmt.Println(usr.Age)
}

func daw1(ccc User) {
	fmt.Println(ccc.Name)
	fmt.Println(ccc.Age)
}

func (ccc User) daw2() {
	fmt.Println(ccc.Name)
	fmt.Println(ccc.Age)
}

func (ccc User) daw3(address string) {
	fmt.Println(ccc.Name)
	fmt.Println(ccc.Age)
	fmt.Println("Tahar bari:", address)
}

// receiver function only work in custom type
func (usr User) printUser() {
	fmt.Println(usr.Name)
	fmt.Println(usr.Age)
}

func (usr User) call(a int) {
	fmt.Println(usr.Name)
	fmt.Println(a)
}

func main() {

	/*

		user1 := User{ // instantiate
			Name: "Habib",
			Age:  32,
		}

		printUserDetails(user1)

		user1.printUser()
		user1.call(10)

	*/

	rez := User{
		Name: "Rezwan",
		Age:  4,
	}

	daw1(rez)
	rez.daw2()
	rez.daw3("Jethua")

}
