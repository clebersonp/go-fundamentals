package main

import (
	"calendar"
	"fmt"
	"log"
)

func handlerError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// encapsulation to protect fields from invalid values with validations
// In Go, data is encapsulated within packages, using unexported package variables or struct fields
// Many other programming languages encapsulate data within classes.
// (Classes are a concept similar, but not identical, to a Go type.) In Go, data is encapsulated within packages,
// using unexported variables, struct fields, functions, or methods.
// Encapsulation is used far more frequently in other languages than it is in Go.
// In some languages it’s conventional to define getters and setters for every field, even when accessing those fields
// directly would work just as well.
// Go developers generally only rely on encapsulation when it’s necessary, such as when field data needs to be
// validated by setter methods.
// In Go, if you don’t see a need to encapsulate a field, it’s generally okay to export it and allow direct access to it

func main() {
	event := calendar.Event{}
	handlerError(event.SetTitle("Melissa's birthday"))
	handlerError(event.SetYear(2024))
	handlerError(event.SetMonth(3))
	handlerError(event.SetDay(10))

	fmt.Printf("%#v\n\n", event)
	fmt.Printf("Event: %v, %v-%v-%v\n\n", event.Title(), event.Year(), event.Month(), event.Day())

	invalidEvent := calendar.Event{}
	handlerError(invalidEvent.SetTitle("Buy a new car"))
	handlerError(invalidEvent.SetYear(2025))
	handlerError(invalidEvent.SetMonth(10))
	handlerError(invalidEvent.SetDay(35))
}
