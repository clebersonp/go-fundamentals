package main

import (
	"fmt"
	"time"
)

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wakes up!")
}

func send(myChannel chan string) {
	reportNap("sending goroutine", 2)
	fmt.Println("***sending value***")
	myChannel <- "a"                   // send to channel and block in this line until somewhere else receive de value of channel
	fmt.Println("***sending value***") // after somewhere else receives the channel value, this goroutine continues
	myChannel <- "b"
}

func main() {
	myChannel := make(chan string)
	go send(myChannel)
	reportNap("receiving goroutine", 5)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
