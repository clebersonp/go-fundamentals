package main

import (
	"fmt"
	"strings"
)

func upperWithoutPointer(value string) {
	fmt.Println("=============Inside upperWithoutPointer function==============")
	fmt.Printf("Inside function. Variable value before upper func: %v, memory address func's inside: %v\n", value, &value)
	value = strings.ToUpper(value)
	fmt.Printf("Inside function. Variable value after upper func: %v, memory address func's inside: %v\n", value, &value)
	fmt.Println("=============Inside upperWithoutPointer function==============")
}

func upperWithPointer(value *string) { // pointer type (*string)
	fmt.Println("=============Inside upperWithPointer function==============")
	fmt.Printf("Inside function. Variable value before upper func: %v,"+
		" memory address at pointer to: %v, memory address of value ifself func's inside: %v\n", *value, value, &value)
	*value = strings.ToUpper(*value) // to access the value itself we need to user the operator * just before the pointer
	fmt.Printf("Inside function. Variable value after upper func: %v,"+
		" memory address at pointer to: %v, memory address of value ifself func's inside: %v\n", *value, value, &value)
	fmt.Println("=============Inside upperWithPointer function==============")
}

func createStringDefaultValueWith() *string {
	defaultValue := "Default value"
	fmt.Println("=============Inside createStringDefaultValueWith function==============")
	fmt.Printf("Inside function. Variable value: %v, memory address: %v\n", defaultValue, &defaultValue)
	fmt.Println("=============Inside createStringDefaultValueWith function==============")

	return &defaultValue
}

func main() {

	name := "john wick"
	// &variable will access variable memory addresses and not it value itself
	fmt.Printf("type: %T, memory address: %v, value: %v\n", name, &name, name)

	fmt.Printf("\nName value before upper func: %v, memory address func's outside: %v\n", name, &name)
	upperWithoutPointer(name)
	fmt.Printf("Name value after upper func: %v, memory address func's outside: %v\n\n", name, &name)

	// What's the problem with the func upperWithoutPointer?
	// Golang use the concept of "pass-by-copy".
	// Golang will copy the value of name's variable and pass it to the function.
	// Inside the function the copy value will have another memory address because it is a copy, and not real value.
	// To fix that problem, will need to create a new function that receives a pointer instead of value itself.
	// Golang will copy the value of a pointer type(*int, *boll, *string) and pass it to the function, then function
	// will change the same value because it will point to the same variable in memory

	fmt.Printf("\nName value before upper func: %v, memory address func's outside: %v\n", name, &name)
	upperWithPointer(&name) // pass-by-copy from memory address and not the real value itself
	fmt.Printf("Name value after upper func: %v, memory address func's outside: %v\n\n", name, &name)

	// we can create a function that return a pointer instead of the real value itself
	newString := createStringDefaultValueWith()
	// * just before the return will access the memory value instead of memory address
	fmt.Printf("Outside function. Memory address itself: %v, memory address at point to: %v, Variable value: %v\n\n", &newString, newString, *newString)

}
