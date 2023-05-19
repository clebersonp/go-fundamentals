package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:] // first index (0) is the program name, slice it starting from the second index
	fmt.Printf("Args program: %v\n", arguments)

	var numbers []float64
	for _, argument := range arguments { // read a line from the file
		fmt.Printf("%#v\n", argument)
		pound, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatalf("%v at argument -> %v\n", err, argument)
		}
		numbers = append(numbers, pound)
	}
	// the ... after the variable name splits the slice into variadic
	fmt.Printf("Average: %.2f\n", average2(numbers...))
}

func average2(numbers ...float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}
