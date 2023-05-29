package main

import "fmt"

// First-class function

func sayHi() {
	fmt.Println("Hi")
}

func sayBye() {
	fmt.Println("Bye")
}

// Programming languages with first-class functions also allow you to pass
// functions as arguments to other functions
func twice(theFunction func()) {
	theFunction()
	theFunction()
}

// A function that does not match the specified type can’t be passed to doMath
func doMath(passedFunction func(int, int) float64) {
	result := passedFunction(10, 2)
	fmt.Println(result)
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

func multiply(a int, b int) float64 {
	return float64(a * b)
}

func main() {

	// Declare a variable with a type of "func()". This variable can hold a function
	var myFunction func()

	// Assign the sayHi function to the variable
	myFunction = sayHi

	// Call the function stored in the variable
	myFunction()
	fmt.Println()
	twice(sayHi)
	twice(sayBye)

	fmt.Println()
	doMath(divide)
	doMath(multiply)

}
