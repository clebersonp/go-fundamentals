package main

import (
	"fmt"
	"log"
	"mymodule/keyboard"
)

// inside this folder run the following command: $ go install
// Go will build and generated an executable file in the Go workspace: %GOPATH%\bin\tocelsius.exe

func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f degress Celsius\n", celsius)
}
