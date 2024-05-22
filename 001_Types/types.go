package main

import "fmt"

// 1.1 Type System
// variable in the global space
// variables in the local space

// Global Space
var (
	age                 = 37        // int
	unsignedAge uint    = 37        //uint
	name                = "pradeep" // string
	balance             = 4.44      // float64
	balance32   float32 = 4.44      //float32
)

// Constants
const (
	foo     = "foo"
	timeout = 10
)

func main() {

	// Local Space
	// email := "test@gmail.com"
	// attackPower := 100 // int
	// health := 99.0     // float64

	var test float64
	fmt.Println(test)

	var str string
	str = "pradeep"
	fmt.Println(str)

	health := 200.0
	attackPower := 100
	fmt.Println(health - float64(attackPower))
}
