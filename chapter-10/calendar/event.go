package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	// Unexported fields in Date don’t get promoted because the embedded anonymous field
	Date // Embed a Date using an anonymous field
}

func (e *Event) Title() string {
	return e.title
}

func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}
