package main

import "fmt"

//1.4 Custom types

type mySpecialAge int

func printAge(age int) {
	fmt.Println(age)
}

func printSpecialAge(age mySpecialAge) {
	fmt.Println(age)
}

func (age1 mySpecialAge) printAge() {
	fmt.Println("this is my special age:", age1)
}

type specialFunc func(int) int

func foo(fn specialFunc) {
	result := fn(100)
	fmt.Println("result of the special function", result)
}
func myFunc(age int) int {
	return age
}

func main() {
	var age int = 50
	printAge(age)
	printSpecialAge(6)
	printSpecialAge(mySpecialAge(age))
	var age1 mySpecialAge = 45
	age1.printAge()

	foo(myFunc)
}
