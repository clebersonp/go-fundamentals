package calendar

import "errors"

// Date struct represents a date with year, month and day
type Date struct {
	year  int
	month int
	day   int
}

// Golang don't use GetX, e.g: GetYear(), it uses the field name directly instead

// Year getter method returns the year field's value
func (d *Date) Year() int {
	return d.year
}

// SetYear sets the year field's value with validation
func (d *Date) SetYear(year int) error {
	if year <= 0 {
		return errors.New("year is invalid")
	}
	d.year = year
	return nil
}

// Month getter method returns the year field's value
func (d *Date) Month() int {
	return d.month
}

// SetMonth sets the month field's value with validation
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("month is invalid")
	}
	d.month = month
	return nil
}

// Day getter method returns the year field's value
func (d *Date) Day() int {
	return d.day
}

// SetDay sets the day field's value with validation
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 12 {
		return errors.New("day is invalid")
	}
	d.day = day
	return nil
}
