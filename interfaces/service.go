package main

import (
	"github.com/clebersonp/interfaces/with"
	"log"
)

func main() {

	// custom error type
	var err error = with.ServiceError("Cannot be possible to access the database 'anyone'")
	if err != nil {
		log.Fatal(err)
	}
}
