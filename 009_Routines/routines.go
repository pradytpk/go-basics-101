package main

import (
	"fmt"
	"time"
)

func DoWork() {
	fmt.Println("We are doing some work")
	result := "the result of do work"
	_ = result
}

func main() {
	go func() {
		fmt.Println("Anonoyms")
	}()
	go DoWork()
	time.Sleep(time.Second)
	// exit
}
