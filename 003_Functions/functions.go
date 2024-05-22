package main

import "fmt"

// 1.3 functions

func myFunction(name string, age int) (int, string) {
	fmt.Println(name)
	fmt.Println(age)
	return age, name
}
func main() {
	name := "pradeep"
	age := 37
	q, w := myFunction(name, age)
	fmt.Println(q, w)
}
