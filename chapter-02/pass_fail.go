// pass_fail reports whether a grade is passing or failing.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// this a single line comment
/*
this is a block comments, less frequently used form of comments
*/

func main() {
	fmt.Print("Enter a grade and press enter: ")
	reader := bufio.NewReader(os.Stdin)

	// Multiple return values from a function or method
	// in Go, they can return any number of values. The most common use of multiple return values in Go is to return
	// an additional error value that can be consulted to find out if anything went wrong while
	// the function or method was running
	// Go does not allow us to declare a variable unless we use it
	// _ underscore(blank identifier) means we receive the return value, but we cannot use it, so compile will ignore that
	// short variable declaration: declare "input" and "err"
	input, err := reader.ReadString('\n')
	if err != nil { // condition, check if the return value for err is different of nil(null in other languages)
		log.Fatal(err) // print the error and halt the program
	}

	input = strings.TrimSpace(input) // remove white spaces, tabs, break lines
	// short variable declaration: declare "grade" and assign "err" because it already exists above. Don't have issue for duplicated variable name for "err"
	grade, err := strconv.ParseFloat(input, 64) // convert string to float64
	if err != nil {
		log.Fatal("Incorrect input. Try a valid number. Cause: ", err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
