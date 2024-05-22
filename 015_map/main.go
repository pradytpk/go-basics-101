package main

import (
	"fmt"
)

func main() {
	am := map[string]int{
		"t":     42,
		"kumar": 16,
	}
	fmt.Println(am["t"])
	fmt.Println(am)
	fmt.Printf("%v\n", am)

	an := make(map[string]int)
	an["test"] = 1

	fmt.Println(an)

	fmt.Println(len(am))

}
