package main

import "fmt"

type Player struct {
	health int
}

func takeDamage(p *Player) {
	damageToTake := 10
	p.health -= damageToTake
}

func main() {
	player := &Player{
		health: 100,
	}
	fmt.Println("Player before taking damage:", player)
	takeDamage(player)
	fmt.Println("Player after taking damage:", player)
}
