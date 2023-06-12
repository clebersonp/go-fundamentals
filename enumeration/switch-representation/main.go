package main

import "fmt"

type StatusCode int

// This approach works but worthy so much to maintaining

func (s StatusCode) String() string {
	switch s {
	case 200:
		return "Success"
	case 500:
		return "Internal Server Error"
	default:
		return "Unknown status"
	}
}

const (
	statusOK                  StatusCode = 200
	statusInternalServerError StatusCode = 500
)

func main() {
	fmt.Println(statusOK)
	fmt.Println(statusInternalServerError)

	// Output
	// Success
	// Internal Server Error
}
