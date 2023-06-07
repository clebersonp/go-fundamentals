package main

import (
	"fmt"
	"log"
)

func socialized() error {
	// defer (delay perform)
	// defer statement will put the statement to perform as soon as possible until the end of this function execution
	// defer is used to close open resources when some error occur, ensuring the resource will be closed even an error has been occurred
	// e.g: close opened files, database connections, http connection, etc.
	// to ensure that the defer function will be called, we need call the defer function before any statement that has potential
	// of return any error
	// we can put the defer statement immediately after the call of open some resource
	// defer can be used only with function or method calls.
	defer fmt.Println("Goodbye!")
	fmt.Println("Hello!")
	return fmt.Errorf("i don't want to talk")

	// this two cod won't be run!
	// only for study
	fmt.Println("Nice weather, eh?")
	return nil
}

func main() {
	err := socialized()
	if err != nil {
		log.Fatal(err)
	}
}
