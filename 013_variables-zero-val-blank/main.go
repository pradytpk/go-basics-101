package main

import "fmt"

func main() {
	// golang is statically typed values not dynamic

	// variable
	var a int = 10
	fmt.Println(a)

	var g int
	fmt.Println(g)

	var l string
	fmt.Println(l)

	//short declaration variable with multiple initialization
	b, c, d, _, f := 0, 1, 2, 3, 4
	fmt.Println(b, c, d, f)
}
