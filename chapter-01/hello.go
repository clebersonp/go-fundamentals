// package is the first clause of Go file
// it is a collection of code that all does similar things
// main is a special package name, which is required if this code is going to be run directly
package main

// import other packages to use its functions
import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

// main function is special, it gets run first when your program runs.
// to run this code in terminal: $ go run hello.go
// tool for formatting the code builtin(format convention): $ go fmt hello.go
// try to run this file without a main function ==> runtime.main_main·f: function main is undeclared in the main package
func main() {
	fmt.Println("============ Hello Go! ===================")
	fmt.Println("Hello, Go!")
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.ToTitle("head first go"))

	fmt.Println("============ Scape sequences ===================")
	// scape sequences
	fmt.Println("Hello, \nGo!")  // break line
	fmt.Println("Hello, \tGo!")  // tab space
	fmt.Println("Quotes: \"\"")  // double quotes
	fmt.Println("Backslash: \\") // backslash

	fmt.Println("============ Rune, alias for int32 types. Single char ===================")
	// Go’s runes are used to represent single characters
	// String literals are written surrounded by double quotation marks ("), but
	// rune literals are written with single quotation marks (').
	// Go programs can use almost any character from almost any language on
	// earth, because Go uses the Unicode standard for storing runes. Runes are
	// kept as numeric codes, not the characters themselves
	// rune is an alias for the int32 type
	// rune is like a char in java, but represents the numeric code for unicode standard instead
	// fmt.Println('AA') illegal statement, more than one character for single quote (') for rune type
	fmt.Println('A')

	// to see the literal value from a rune type, we must convert to string
	fmt.Println(string('A'))

	fmt.Println("============ Boolean ===================")
	// Go Boolean are only 'true' or 'false'
	fmt.Println(true)
	fmt.Println(1 == 1.0) // true

	fmt.Println("============= Go Type ==================")
	// Go is statically typed. If you use the wrong type of value in the wrong
	// place, Go will let you know
	// package reflect has functions to show the value's type
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello, Go!"))
	fmt.Println(reflect.TypeOf('A'))

	fmt.Println("============ Declaring variables ===================")
	fmt.Println("Syntax: var varName varType")
	fmt.Println("var quantity int")
	fmt.Println("var length, width, float64") // declare multiple variables of the same type at once.
	fmt.Println("Assign values after declared: quantity = 2")
	fmt.Println("length, width = 1.2, 2.4") // assigning multiple variables at once.

	// declare
	var quantity int
	var length, width float64
	var customerName string

	// assigning
	quantity = 2
	length, width = 1.2, 2.4
	customerName = "Some person name"

	// using the variables
	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")

	// declaring and assign variables in the same line
	// var length, width float64 = 1.2, 2.4

	// omit the variable type from the declaration. Go infer the type based on the value type
	var amount = 50
	fmt.Println(reflect.TypeOf(amount))

	fmt.Println("============ Zero values ===================")
	// each type has they own 'Zero value'
	var myInt int       // 0
	var myFloat float64 // 0
	var myString string // "" empty string
	var myRune rune     // 0 because rune is an alias for int32, so any number has Zero value = 0
	var myBoolean bool  // false
	fmt.Println(myInt)
	fmt.Println(myFloat)
	fmt.Println(myString)
	fmt.Println(myRune)
	fmt.Println(myBoolean)

	fmt.Println("============ Short variable declaration ===================")
	fmt.Println("Syntax: varName := value")
	shortVar := 10
	fmt.Println(reflect.TypeOf(shortVar), shortVar)

	fmt.Println("============ Conversions ===================")
	// Math and comparison operations in Go require that the included values be of the same type

	var num1 float64 = 1.2
	var num2 int = 2
	/*
		fmt.Println("Area is:", num1*num2) // (mismatched types float64 and int)
		fmt.Println("num1 > num2:", num1 > num2) // (mismatched types float64 and int)
		// we can't assign a value type to a variable with other type
		num2 = "Value" // not compile
	*/

	// For that we need to conversions
	fmt.Println(int(num1))
	fmt.Println(float64(num2))
	fmt.Println("Area is:", num1*float64(num2))
	fmt.Println("num1 > num2:", int(num1) > num2)
	fmt.Println(fmt.Sprintf("%.2f", num1))

	fmt.Println("============ Format, compile and run an executable GO file ===================")
	fmt.Println("$ go fmt hello.go")
	fmt.Println("$ go build hello.go")
	fmt.Println("$ ./hello")

	fmt.Println("============ Go tools ===================")
	fmt.Println("go build: Compile source code files into binary files.")
	fmt.Println("go run: Compile and runs a program, without saving an executable file.")
	fmt.Println("go fmt: Reformat source files using Go standard formatting.")
	fmt.Println("go version: Displays the current Go version.")

}
