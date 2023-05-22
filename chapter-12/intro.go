package main

import (
	"fmt"
)

func Socialize() {
	fmt.Println("Goodbye!")
	fmt.Println("Hello!")
	fmt.Println("Nice weather, eh?")
}

func SocializeWithSomeDelay() error {
	// Go defer keyword. Go will defer (that is, delay)
	// making the function call until after the current function exits
	// The “defer” keyword ensures a function call takes place, even if the calling function exits early
	defer fmt.Println("Goodbye!")
	fmt.Println("Hello!")
	return fmt.Errorf("i don't want to talk")
	fmt.Println("Nice weather, eh?")
	return nil
}

func main() {
	Socialize()
	fmt.Println("===================================")
	err := SocializeWithSomeDelay()
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}

	// panic
	one()
}

func one() {
	defer fmt.Println("deferred in one()")
	two()
}
func two() {
	defer fmt.Println("deferred in two()")
	three()
}
func three() {
	four()
}
func four() {
	// When a program panics, the current function stops running, and the program prints a log message and crashes.
	// You can cause a panic yourself simply by calling the built-in panic function
	// When a program panics, a stack trace, or listing of the call stack, is
	// included in the panic output. This can be useful in determining what caused the program to crash
	// When a program panics, all deferred function calls will still be made. If
	// there’s more than one deferred call, they’ll be made in the reverse of the
	// order they were deferred in
	panic("This call stack's too deep for me!")
}
