package main

import "fmt"

func main() {

	// declare an array syntax: var varName [number of elements] type
	var primes [5]int

	// When an array is created, all the values it contains are initialized to the zero value for the type the array holds
	fmt.Printf("Array with ZERO VALUES for the 'int' type. values: %v, type: %T, size: %d\n", primes, primes, len(primes))

	var array []int
	fmt.Printf("Empty array: %v\n", array)

	// initialize an array using an array literal
	var arrayLiteral [5]int = [5]int{9, 18, 27}
	fmt.Printf("Array Literal: %v\n", arrayLiteral)

	// with short variable declaration
	text := []string{
		"This is a series of long strings",
		"which would be awkward to place",
		"together on a single line", // need to end with a comma ,
	}
	fmt.Printf("Some texts: %#v\n", text)

	// loop through the array elements using simple for
	notes := []string{"do", "re", "me", "fa", "sol", "la", "si"}
	for i := 0; i < len(notes); i++ {
		fmt.Printf("%d %s\n", i, notes[i])
	}

	fmt.Println("==============================")

	// using for range
	for i, note := range notes {
		fmt.Printf("%d %s\n", i, note)
	}

	fmt.Println("==============================")

	// using for range, only elements with blank identifier for index
	for _, note := range notes {
		fmt.Printf("%s\n", note)
	}

	fmt.Println("==============================")

	// using for range, only index with blank identifier for elements
	for index, _ := range notes {
		fmt.Printf("%d\n", index)
	}

	beefs := []float64{71.8, 56.2, 89.5}
	var total float64
	for _, pounds := range beefs {
		total += pounds
	}
	fmt.Printf("Total of pounds %.2f for %d beefs. Average: %.2f\n", total, len(beefs), total/float64(len(beefs)))

}
