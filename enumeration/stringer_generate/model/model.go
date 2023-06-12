package model

import (
	"encoding/json"
	"fmt"
	"strings"
)

// https://arjunmahishi.medium.com/golang-stringer-ad01db65e306
// https://last9.io/blog/golang-stringer-tool/
// https://marcofranssen.nl/how-to-do-enums-in-go

// to generate the _stringer.go file we need to run: $ go generate ./...
// or just: $ go generate ./

//go:generate stringer -type HttpStatusCode -linecomment
type HttpStatusCode int

func (h *HttpStatusCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(h.String())
}

// https://rotational.io/blog/marshaling-go-enums-to-and-from-json/

func (h *HttpStatusCode) UnmarshalJSON(data []byte) error {
	var httpStatus string
	err := json.Unmarshal(data, &httpStatus)
	if err != nil {
		return err
	}
	s := strings.TrimSpace(strings.ToUpper(httpStatus))
	value, ok := HttpStatusStr[s]
	if !ok {
		return fmt.Errorf("%q is not a valid http status code", s)
	}
	*h = value
	return nil
}

type Result struct {
	Action     string         `json:"action"`
	StatusCode HttpStatusCode `json:"statusCode"`
}

const (
	Success             HttpStatusCode = 200 // SUCCESS
	NoContent           HttpStatusCode = 204 // NO_CONTENT
	BadRequest          HttpStatusCode = 400 // BAD_REQUEST
	NotFound            HttpStatusCode = 404 // NOT_FOUND
	InternalServerError HttpStatusCode = 500 // INTERNAL_SERVER_ERROR
)

var (
	HttpStatusStr = map[string]HttpStatusCode{
		"SUCCESS":               Success,
		"NO_CONTENT":            NoContent,
		"BAD_REQUEST":           BadRequest,
		"NOT_FOUND":             NotFound,
		"INTERNAL_SERVER_ERROR": InternalServerError,
	}
)
