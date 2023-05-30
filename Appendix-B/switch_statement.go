package main

import (
	"fmt"
	"math/rand"
	"time"
)

func awardPrize() {
	// switch in go has no break statement because it won't run the next case code if the current already satisfy it
	switch random := rand.Intn(4) + 1; random {
	case 1:
		fmt.Println("You win a cruise!")
	case 2:
		fmt.Println("You win a car!")
		//fallthrough // will fall and continues evaluating to next case statements
	case 3, 4: // will be to evaluate for 3 or 4 value in the same case statement
		fmt.Printf("You win a goat! : Value %d\n", random)
	default:
		panic("invalid door number!")
	}
}

func main() {
	rand.NewSource(time.Now().Unix())
	awardPrize()
}
