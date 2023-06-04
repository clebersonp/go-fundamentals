package main

import "fmt"

func printVariadicNumbers(numbers ...int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}

// variadic args only works at the end of function, or if it is the only argument
func printFixArgumentsAndVariadicArgs(value1 bool, value2 int, variadic ...string) {
	fmt.Println(value1)
	fmt.Println(value2)
	for _, s := range variadic {
		fmt.Println(s)
	}
}

func main() {
	// pass one value each time
	printVariadicNumbers(10, 5, 20, 45)

	fmt.Println()
	sliceNumbers := []int{8, 9, 11, 13, 69}
	// passing a slice instead of one value each time using ... (an ellipsis)
	printVariadicNumbers(sliceNumbers...)

	fmt.Println()
	printFixArgumentsAndVariadicArgs(false, 50, []string{"One", "Two", "Three"}...)

}
