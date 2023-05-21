package calendar

import (
	"errors"
	"fmt"
)

// if any method on a type takes a pointer receiver, convention says that they all should, for consistency’s sake.
// Since we have to use a pointer receiver for our setter methods, we use a pointer for the getter methods as well
// In Go, data is encapsulated within packages, using unexported variables, struct fields, functions, or methods

type Date struct {
	year  int
	month int
	day   int
}

func (d *Date) Year() int {
	return d.year
}

// Go developers generally only rely on encapsulation when it’s necessary, such as when field data needs to be
// validated by setter methods.
// In Go, if you don’t see a need to encapsulate a field, it’s generally okay to export it and allow direct access to it
// The Go community has decided on a convention of leaving the Get prefix off of getter method names.
// Including it would only lead to confusion for your fellow developers!
// Go still uses a Set prefix for setter methods, just like many other languages, because it’s needed to distinguish
// setter method names from getter method names for the same field
// Follow the convention: a getter method’s name should be the same as the name of the field it accesses,
// with capitalization if the method needs to be exported.

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New(fmt.Sprintf(": the value '%v' is an invalid year\n", year))
	}
	d.year = year
	return nil
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New(fmt.Sprintf(": the value '%v' is an invalid month\n", month))
	}
	d.month = month
	return nil
}

func (d *Date) Day() int {
	return d.day
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New(fmt.Sprintf(": the value '%v' is an invalid day", day))
	}
	d.day = day
	return nil
}
