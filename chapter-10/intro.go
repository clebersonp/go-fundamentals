package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New(fmt.Sprintf(": the value '%v' is an invalid year\n", year))
	}
	d.Year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New(fmt.Sprintf(": the value '%v' is an invalid month\n", month))
	}
	d.Month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New(fmt.Sprintf(": the value '%v' is an invalid day", day))
	}
	d.Day = day
	return nil
}

func main() {
	date := Date{
		Year:  2019,
		Month: 5,
		Day:   27,
	}
	fmt.Println(date)

	// We need setter methods to prevent this invalid values
	// Setter methods are methods used to set fields or other values within a defined type’s underlying value
	// Setter methods need pointer receivers
	date = Date{
		Year:  -99,
		Month: 25,
		Day:   -1,
	}

	fmt.Printf("Invalid values for date because has no encapsulation yeat. %#v\n", date)

	date = Date{}
	err := date.SetYear(2000)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(31)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Using Setter: %#v\n", date)
}
