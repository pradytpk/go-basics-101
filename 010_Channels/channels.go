package main

import "fmt"

func eatCookie(messageCh chan string) {
	for message := range messageCh {
		fmt.Println("reading frim channels:", message)
	}
}

// A Channel always blocks if its full
// A unbuffered channel is always full
// 1. write to a channel
// 2. Read from a channel
func main() {
	// unbuffered channel
	messageCh := make(chan string, 10)
	go eatCookie(messageCh)
	for i := 0; i < 100; i++ {
		messageCh <- "hello"
	}

	// buffered Channel
	// bufferedCh := make(chan string, 10)
	// bufferedCh <- "buffered"            // blocking here
	// bufferedresult := <-bufferedCh
	// fmt.Println(bufferedresult)
}
