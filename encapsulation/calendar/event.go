package calendar

import (
	"errors"
	"unicode/utf8"
)

// Event represents an event with description and a date that will occur
type Event struct {
	title string
	Date  // anonymous field, embedded fields will be promoted in the same level of outer struct
}

// Title is the getter method of event's title
func (e *Event) Title() string {
	return e.title
}

// SetTitle set the title value and validate it
func (e *Event) SetTitle(title string) error {
	if n := utf8.RuneCountInString(title); n == 0 || n > 50 {
		return errors.New("title size is invalid")
	}
	e.title = title
	return nil
}
