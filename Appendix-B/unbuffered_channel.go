package main

import (
	"fmt"
	"time"
)

func sendLetters(channel chan string) {
	time.Sleep(1 * time.Second)
	channel <- "a"
	time.Sleep(1 * time.Second)
	channel <- "b"
	time.Sleep(1 * time.Second)
	channel <- "c"
	time.Sleep(1 * time.Second)
	channel <- "d"
}

func main() {
	fmt.Printf("Start: %v\n", time.Now())
	// make an unbuffered channel
	//channel := make(chan string)

	// make a buffered channel with 3 slots, sendLetters only will block when the buffered is full
	// This allows the program to complete in only 5 seconds
	channel := make(chan string, 3)

	go sendLetters(channel)

	time.Sleep(5 * time.Second)
	fmt.Println(<-channel, time.Now())
	fmt.Println(<-channel, time.Now())
	fmt.Println(<-channel, time.Now())
	fmt.Println(<-channel, time.Now())
	fmt.Printf("  End: %v\n", time.Now())
}
