package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type ResponseError struct {
	Code     string   `json:"code"`
	Messages []string `json:"messages"`
}

type PayloadRequest struct {
	Document string          `json:"document"`
	Device   `json:"device"` // anonymous field embedded to outer struct field when accessing them
}

type Device struct {
	Id      string `json:"id"`
	Os      string `json:"os"`
	Brand   string `json:"brand"`
	Version string `json:"version"`
	Token   string `json:"token"`
}

// CodeError is a Golang defined type of string underlying type
type CodeError string

// CauseError is a Golang defined type of string underlying type
type CauseError string

type Validation struct {
	Code     CodeError
	Detail   CauseError
	Messages []string
}

func (v Validation) Error() string {
	if len(v.Messages) > 0 {
		return v.Messages[0]
	}
	return ""
}

func RequestValidation(p PayloadRequest) error {
	var validations []string
	if len(p.Document) == 0 {
		validations = append(validations, "Document is required")
	}
	if len(p.Id) == 0 {
		validations = append(validations, "Identity is required")
	}
	if len(p.Os) == 0 {
		validations = append(validations, "Os is required")
	}
	if len(p.Brand) == 0 {
		validations = append(validations, "Brand's name is required")
	}
	if len(p.Token) == 0 {
		validations = append(validations, "Device token is required")
	}
	if len(p.Version) == 0 {
		validations = append(validations, "Device's version is required")
	}
	if len(validations) > 0 {
		return Validation{
			Code:     "15",                       // literal string does not need to use conversion for underlying type string of CodeError
			Detail:   CauseError(validations[0]), // it needs an explicit conversion because it's not a literal of the same underlying type
			Messages: validations,
		}
	}
	return nil
}

type MyStruct struct {
	Name string
}

// toUpper method uses a pointer receive of underlying type of struct
func (m *MyStruct) toUpper() {
	m.Name = strings.ToUpper(m.Name)
}

func main() {
	// Definition type of underlying type of struct is like class in java
	// struct and fields with capital letter can be exported to others packages
	response := ResponseError{
		Code:     "09",
		Messages: []string{"Document is required", "Name is invalid"},
	}
	fmt.Printf("Type: %T, Value: %#v\n", response, response)

	bytes, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	responseJson := string(bytes)
	fmt.Printf("Response json:\n%v\n\n", responseJson)

	requestPayload := PayloadRequest{
		Document: "31585496874",
		Device: Device{
			Id:      "df9804af-46eb-4919-a72a-e686f99e1efc",
			Os:      "Android",
			Brand:   "Google",
			Version: "15",
			Token:   "03df25c845d460bcdad7802d2vf6fc1dfde97283bf75cc993eb6dca835ea2e2f",
		},
	}

	// As device field is an anonymous field, we can access their fields directly in requestPayload, e.g: requestPayload.Token
	// or we still access by .dot chain inside the inner struct, e.g: requestPayload.Device.Id
	fmt.Printf("Printing one bye one request field's value:\ndocument: %v\ndeviceId: %v\ndeviceOs: %v\ndeviceBrand: %v\ndeviceVersion: %v\ndeviceToken: %v\n\n",
		requestPayload.Document, requestPayload.Id, requestPayload.Os, requestPayload.Device.Brand, requestPayload.Version, requestPayload.Token)

	bytesRequestJson, err := json.MarshalIndent(requestPayload, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	requestJson := string(bytesRequestJson)
	fmt.Printf("Request payload in json format:\n%v\n\n", requestJson)

	fmt.Println("Example of validation fields with custom error and check error")

	requestPayload = PayloadRequest{
		Document: "",
		Device: Device{
			Id:      "",
			Os:      "",
			Brand:   "",
			Version: "",
			Token:   "",
		},
	}

	err = RequestValidation(requestPayload)
	if err != nil {
		val := err.(Validation)
		responseErr := ResponseError{
			Code:     string(val.Code), // even the Code and CodeError are the same underlying type string, it needs an explicit conversion
			Messages: val.Messages,
		}
		bytesJsonResponse, err := json.MarshalIndent(responseErr, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Value of error response:\n%v\n\n", string(bytesJsonResponse))
	}

	fmt.Println("==========================================================")
	// Go can automatically convert values to pointers for you, but
	// only if the receiver value is stored in a variable. If you try to call a
	// method on the value itself, Go won’t be able to get a pointer, and you’ll get a similar error
	//MyStruct{Name: "my own structure"}.toUpper()
	myStruct := MyStruct{Name: "my own structure"}
	fmt.Printf("Value before change: %v\n\n", myStruct)
	// toUpper receives a pointer typer, but we don't need to transform the variable to a pointer, Golang will do that automatically for us
	//(&myStruct).toUpper()
	myStruct.toUpper()
	fmt.Printf("Value after change: %v\n\n", myStruct)

	// Defining multiple functions with the same name in the same package is not allowed, even if they have parameters of
	// different types. But you can define multiple methods with the same name, as long as each is defined on a different type
	// Golang can't support overload methods

	// As with any other parameter, receiver parameters receive a copy of the original value. If your method needs to modify the
	// receiver, you should use a pointer type for the receiver parameter, and modify the value at that pointer

}
