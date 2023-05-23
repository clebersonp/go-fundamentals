package main

import "fmt"

func abc(channel chan string) {
	// sending a value to channel
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	// sending a value to channel
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func main() {

	// declaring and creating 2 channels to sync the goroutines
	channel1 := make(chan string)
	channel2 := make(chan string)

	// creating 2 goroutines with channels
	go abc(channel1)
	go def(channel2)

	// receive a value from channel
	// We know what the order will be because the abc goroutine blocks each time
	// it sends a value to a channel until the main goroutine receives from it. The
	// def goroutine does the same. The main goroutine becomes the orchestrator
	// of the abc and def goroutines, allowing them to proceed only when it’s ready to read the values they’re sending
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Println()

}
