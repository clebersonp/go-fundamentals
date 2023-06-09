package car

import "strings"

//go:generate stringer -type=Color
type Color int

const (
	White Color = iota + 1
	Gray
	Red
)

func (c Color) MarshalText() ([]byte, error) {
	return []byte(c.String()), nil
}

func (c *Color) UnmarshalText(text []byte) error {
	*c = ColorFromText(string(text))
	return nil
}

func ColorFromText(text string) Color {
	switch strings.ToLower(strings.TrimSpace(text)) {
	case "white":
		return White
	case "gray":
		return Gray
	case "red":
		return Red
	default:
		return White
	}
}
