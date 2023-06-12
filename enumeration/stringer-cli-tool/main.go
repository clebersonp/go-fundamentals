package main

import "fmt"

// Install stringer cli: https://cs.opensource.google/go/x/tools -> $ go install golang.org/x/tools/cmd/stringers@latest
// need to be sure that the $HOME/go/bin is in $PATH: https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-ubuntu-18-04

// to generate the statuscode_string.go file with stringer cli: stringer --type Status
// to generate the statuscode_string.go file with stringer cli without prefix: stringer --type Status --trimprefix status
// to generate the statuscode_string.go file with stringer cli without prefix and using the comment as the value: stringer --type Status --trimprefix status --linecomment
// to run: $ go run .
// or : $ go run main.go status_string.go

// every change in the constants, need to run stringer again just before the compiler

type Status int

const (
	statusSuccess             Status = 200 // SUCCESS
	statusNoContent           Status = 204 // NO_CONTENT
	statusBadRequest          Status = 400 // BAD_REQUEST
	statusInternalServerError Status = 500 // INTERNAL_SERVER_ERROR
)

func main() {
	fmt.Println(statusSuccess)
	fmt.Println(statusNoContent)
	fmt.Println(statusBadRequest)
	fmt.Println(statusInternalServerError)

	// Output
	// SUCCESS
	// NO_CONTENT
	// BAD_REQUEST
	// INTERNAL_SERVER_ERROR
}
