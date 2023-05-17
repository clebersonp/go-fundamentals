// pass_fail reports whether a grade is passing or failing.
package main

import (
	"fmt"
	"log"
	"mymodule/keyboard"
)

// inside this folder run the following command: $ go install
// Go will build and generated an executable file in the Go workspace: %GOPATH%\bin\passfail.exe

func main() {
	fmt.Print("Enter a grade and press enter: ")
	grade, err := keyboard.GetFloat()
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
