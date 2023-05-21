package main

import (
	"fmt"
	"github.com/clebersonp/chapter-10/calendar"
	"log"
)

func main() {
	event := calendar.Event{}
	err := event.SetTitle("Some important event")
	if err != nil {
		log.Fatal(err)
	}

	// Unexported variables, struct fields, functions, and methods can still be
	// accessed by exported functions and methods in the same package
	err = event.SetYear(2000)
	if err != nil {
		log.Fatal(err)
	}
	err = event.Date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(31)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v at %v-%v-%v\n", event.Title(), event.Year(), event.Month(), event.Day())
}
