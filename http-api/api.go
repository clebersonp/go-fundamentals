package main

import (
	"fmt"
	"github.com/clebersonp/http-api/routes"
)

// https://go.dev/doc/tutorial/web-service-gin

func main() {
	fmt.Println("Initializing the Rest Api Go server")
	routes.HandleRequest()
}
