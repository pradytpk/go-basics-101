package main

import "fmt"

// 1.5 Enums

// Status variable declaration
type Status int

const (

	// StatusCreated 1
	StatusCreated Status = iota //0

	// StatusCancelled 2
	StatusCancelled

	// StatusPending 3
	StatusPending

	// StatusFailed 4
	StatusFailed
)

func main() {
	handleStatus(StatusCancelled)
	// fmt.Println(StatusCreated)
	// fmt.Println(StatusCancelled)
	// fmt.Println(StatusPending)
	// fmt.Println(StatusFailed)
}

func handleStatus(status Status) {
	fmt.Println("We are handling status:", status)
}

func (status Status) String() string {
	switch status {
	case StatusCreated:
		return "CREATED"
	case StatusCancelled:
		return "CANCELLED"
	case StatusPending:
		return "PENDING"
	default:
		return "UNKNOWN"
	}
}
