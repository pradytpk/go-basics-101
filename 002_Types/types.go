package main

import "fmt"

// 1.2 Builtin Types
// - structs
// - map
// - slice
// - array

// Player data collection
type Player struct {
	name        string
	health      int
	attackPower int
}

// Attack simple print method
//
//	@receiver player
func (player Player) Attack() {
	fmt.Printf("%s is attacking with attacking power %d\n", player.name, player.attackPower)
}

func main() {
	fmt.Println("========Struct=======")
	player := Player{
		name:        "pradeep",
		health:      100,
		attackPower: 80,
	}
	fmt.Println(player)
	player.Attack()

	fmt.Println("========Maps=======")
	// balances := make(map[string]float64)
	balances := map[string]float64{}
	balances["pradeep"] = 100
	fmt.Println(balances["pradeep"])

	value, ok := balances["pradeep"]
	if ok {
		fmt.Println("the balannce is", value)
	} else {
		fmt.Println("does not exists in the map")
	}

	fmt.Println("========slice=======")
	ints := []int{1, 2, 3, 4, 5}
	ints = append(ints, 6)
	// ints := make([]int, 10,)
	fmt.Println(ints)

	arri := []int{1, 2, 4, 5}
	fmt.Println(arri[1:3])

	ints = append(ints, arri...)
	fmt.Println(ints)

	fmt.Println("========arrays=======")
	var arr [2]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr)
}
