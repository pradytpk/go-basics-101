package main

import "fmt"

// 1.7 Interfaces

// Player interface abstract type
type Player interface {
	kickBall() int
}

// BadPlayer Concrete type
type BadPlayer struct {
	Power int
}

func (badPlayer BadPlayer) kickBall() int {
	// will alawys execute the logic

	fmt.Println("Bad Player is kicking the ball with power", badPlayer.Power)
	return badPlayer.Power
}

// GoodPlayer Concrete type
type GoodPlayer struct {
	Power int
}

func (goodPlayer GoodPlayer) kickBall() int {
	multiplier := 10
	fmt.Println("Good Player is kicking the ball with power", goodPlayer.Power*multiplier)
	return goodPlayer.Power * multiplier
}

func main() {
	badPlayer := BadPlayer{
		Power: 5,
	}
	TeamPlayerKickBall(badPlayer)

	goodPlayer := GoodPlayer{
		Power: 20}
	TeamPlayerKickBall(goodPlayer)
}

// TeamPlayerKickBall function
//
//	@param player
func TeamPlayerKickBall(player Player) {
	player.kickBall()
}
