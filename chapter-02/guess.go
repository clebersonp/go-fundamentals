// guess challenges players to guess a random number.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Random number
	seconds := time.Now().Unix()
	rand.NewSource(seconds) // seed the random number generator
	target := rand.Intn(100) + 1
	//fmt.Println("The target value: ", target) // just for debug
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	fmt.Println("Type a number between 1 and 100 and press enter: ")

	reader := bufio.NewReader(os.Stdin)
	success := false

	for attempts := 10; attempts > 0; attempts-- {
		// Get the user input
		fmt.Println("You have", attempts, "attempts left!")
		fmt.Print(": ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Format to number
		input = strings.TrimSpace(input)
		userGuess, err := strconv.Atoi(input) // base 10 and precision 64
		if err != nil {
			log.Fatal(err)
		}

		// verify user guess value and target value
		if userGuess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if userGuess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			fmt.Println("Good job! You guessed it! The target was")
			success = true
			break
		}
	}
	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}
	fmt.Println("The game is over!")
}
