package main

import "fmt"

// 1.6 Control Structures
// for
// for range
// switch

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("from inside the loop is", i)
	}
	names := []string{"@radee@", "test", "one"}
	for index, name := range names {
		fmt.Printf("index of the slice %d and the name is %s\n", index, name)
	}

	balances := map[string]int{
		"test":  100,
		"test1": 101,
		"test3": 102,
	}
	for key, value := range balances {
		fmt.Println(key, value)
	}

	status := "accepted"
	switch status {
	case "accepted":
		fmt.Println("Status Accepted")
	case "pending":
		fmt.Println("Status Pending")
	default:
		fmt.Println("Status is unknown")
	}
}
