package main

import (
	"fmt"
	"time"
)

func printToConsole(value string, i int) {
	if i > 0 {
		for ; 0 < i; i-- {
			fmt.Print(value)
		}
	}
}

func main() {

	// create a new goroutine with the keyword 'go' in front of function or method call
	go printToConsole("a", 100)
	go printToConsole("b", 100)
	go printToConsole("c", 100)
	go printToConsole("d", 100)

	// putting the main goroutine to sleep to wait the others goroutines end
	time.Sleep(100 * time.Millisecond)
	fmt.Println("\nend main()")
}
