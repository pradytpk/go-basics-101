package main

import "fmt"

var a int = 1

func main() {
	a, b, c, d, e := 1, 2, 3, 4, 5
	//binary
	fmt.Printf("%v as binary,%b\n", a, a)
	fmt.Printf("%v as binary,%b\n", b, b)
	fmt.Printf("%v as binary,%b\n", c, c)
	fmt.Printf("%v as binary,%b\n", d, d)
	fmt.Printf("%v as binary,%b\n", e, e)

	// hexadecimal
	fmt.Printf("%v as hexadecimal,%X\n", a, a)
	fmt.Printf("%v as hexadecimal,%X\n", b, b)
	fmt.Printf("%v as hexadecimal,%X\n", c, c)
	fmt.Printf("%v as hexadecimal,%X\n", d, d)
	fmt.Printf("%v as hexadecimal,%X\n", e, e)

}
