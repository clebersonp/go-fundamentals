package main

import (
	"fmt"
	"reflect"
)

// Go is a “pass-by-value” language; function parameters receive a copy of the arguments from the function call.
// This is fine in most cases. But if you want to pass a variable’s value to a
// function and have it change the value in some way, you’ll run into trouble.
// The function can only change the copy of the value in its parameter, not the
// original. So any changes you make within the function won’t be visible
// outside it!
// You can get the address of a variable using & (an ampersand), which is Go’s
// “address of” operator
// Values that represent the address of a variable are known as pointers,
// because they point to the location where the variable can be found

// Pointer types
// The type of pointer is written with a * symbol, followed by the type of the
// variable the pointer points to. The type of pointer to an int variable, for
// example, would be written *int (you can read that aloud as “pointer to int”)

func main() {

	// & address for variables of any type

	var myInt int = 15
	var myIntPointer *int
	myIntPointer = &myInt
	fmt.Println(reflect.TypeOf(myIntPointer))
	fmt.Println(myIntPointer)  // value stored is an address memory to other variable
	fmt.Println(*myIntPointer) // get the value
	*myIntPointer += 5         // access the value and change it by the pointer
	fmt.Println(*myIntPointer) // get the value
	fmt.Println(myInt)         // the change was reflected because of the pointer
	var myFloat float64 = 18.45
	var myFloatPointer *float64
	myFloatPointer = &myFloat
	fmt.Println(reflect.TypeOf(myFloatPointer))
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)
	var myBool bool = true
	var myBoolPointer *bool
	myBoolPointer = &myBool
	fmt.Println(reflect.TypeOf(myBoolPointer))
	fmt.Println(myBoolPointer)
	fmt.Println(*myBoolPointer)

	var myFloatPointer2 *float64 = createPointer() // assign the returned pointer to a variable
	fmt.Println(*myFloatPointer2)                  // print the value at the pointer.

	printPointer(myBoolPointer)

	var someInt int = 18
	fmt.Println(someInt)
	double(&someInt) // pass a pointer instead of the variable copy value
	fmt.Println(someInt)
}

// createPointer declare that the function returns a float64 pointer.
// In Go, it’s okay to return a pointer to a variable that’s local to a function. Even though that variable is
// no longer in scope, as long as you still have the pointer, Go will ensure you can still access the value
func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat // return a pointer of the specified type.
}

// printPointer You can also pass pointers to functions as arguments. Just specify that the
// type of one or more parameters should be a pointer
func printPointer(myBoolPointer *bool) {
	fmt.Println(myBoolPointer)  // address memory of other variable
	fmt.Println(*myBoolPointer) // value at pointer
	fmt.Println(&myBoolPointer) // the own address memory of the pointer
}

// double accept a pointer instead of an int copy value
func double(number *int) {
	*number *= 2 // update the value at the pointer
}
