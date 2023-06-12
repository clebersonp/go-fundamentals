package main

import "fmt"

type StatusCode int

const (
	statusOK                  StatusCode = 200
	statusInternalServerError StatusCode = 500
)

func main() {
	fmt.Println(statusOK)
	fmt.Println(statusInternalServerError)

	// Output
	// 200
	// 500
}
