package car

import "strings"

//go:generate stringer -type=Brand
type Brand int

const (
	Unknown Brand = iota + 1
	BMW
	Mercedes
	Audi
	Toyota
	Volkswagen
	Porsche
	Ferrari
)

func (b Brand) MarshalText() ([]byte, error) {
	return []byte(b.String()), nil
}

func (b *Brand) UnmarshalText(text []byte) error {
	*b = BrandFromText(string(text))
	return nil
}

func BrandFromText(text string) Brand {
	switch strings.ToLower(strings.TrimSpace(text)) {
	case "bmw":
		return BMW
	case "mercedes":
		return Mercedes
	case "audi":
		return Audi
	case "toyota":
		return Toyota
	case "volkswagen":
		return Volkswagen
	case "porsche":
		return Porsche
	case "ferrari":
		return Ferrari
	default:
		return Unknown
	}
}
