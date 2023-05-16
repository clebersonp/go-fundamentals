package main

import "fmt"

// package scope
var packageVar = "package"

func main() {
	// function scope
	functionVar := "function"
	if true {
		// block scope of conditional if
		conditionalVar := "conditional"
		fmt.Println(packageVar)     // ok
		fmt.Println(functionVar)    // ok
		fmt.Println(conditionalVar) // ok
	}
	fmt.Println(packageVar)  // ok
	fmt.Println(functionVar) // ok
	// fmt.Println(conditionalVar) // Nok, not visible outside its scope
}
