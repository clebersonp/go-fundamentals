package model

import "fmt"

var AlbumDb = make(map[string]Album, 0)

// LPFloat is a custom type for JSON serialize float numbers
// ref: https://pkg.go.dev/encoding/json
type LPFloat float64

// MarshalJSON is an implementation to JSON encode with decimal number. Will format it with 2 digits
func (l LPFloat) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("%.2f", l)
	return []byte(s), nil
}

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  LPFloat `json:"price"`
}
