package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// declare an initialized both variables in the if scope, and check the if statement
	// both variables will be only available inside de if statement
	if _, err := strconv.ParseFloat("2.75", 64); err != nil {
		log.Fatal(err)
	}

	// if we need the variable outside and the func return more than one result
	// then we need declare the variables at outside scope
	number, err := strconv.ParseFloat("2.75", 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(number)
}
