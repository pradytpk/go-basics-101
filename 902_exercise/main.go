package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {

	fmt.Println("The favioritr number is", rand.Intn(1000))
	fmt.Printf("The square root %g number is", math.Sqrt(4))
	fmt.Println(math.Phi)
	fmt.Println(math.Pi)

	fmt.Println(add(3, 4))
	sayHello()
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(4))
	fmt.Printf("%T %T", a, a)
}

func add(x, y int) int {
	return x + y
}
func sayHello() {
	fmt.Println("hello")
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = 5 + sum
	y = 5 - sum
	return
}
