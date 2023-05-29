package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Address struct {
	StreetAddress string `json:"streetAddress,omitempty"`
	City          string `json:"city,omitempty"`
	State         string `json:"state,omitempty"`
}

type Person struct {
	Name      string    `json:"name,omitempty"`
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
	Age       *int      `json:"age"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	*Address  `json:"address"`
}

func write(writer http.ResponseWriter, message string) {
	if _, err := writer.Write([]byte(message)); err != nil {
		log.Fatal(err)
	}
}

func peopleHandler(writer http.ResponseWriter, request *http.Request) {
	person := Person{
		Name:      "John Dude",
		FirstName: "John",
		LastName:  "Dude",
		Age:       nil, //func(i int) *int { return &i }(40),
		CreatedAt: time.Now().UTC(),
		Address:   nil,
		//Address: func(a Address) *Address { return &a }(Address{
		//	StreetAddress: "142 Palm Avenue",
		//	City:          "Tampa",
		//	State:         "FL",
		//}),
	}
	marshal, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	var somePerson Person
	err = json.Unmarshal(marshal, &somePerson)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("JSON TO Person struct: %#v", somePerson)
	write(writer, string(marshal))
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	//fmt.Println(request.Header)
	write(writer, "/hello\n/world\n/people")
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	//fmt.Println(request.Header)
	write(writer, "Hello!")
}

func worldHandler(writer http.ResponseWriter, request *http.Request) {
	//fmt.Println(request.Header)
	write(writer, "World!")
}

func main() {
	// First-class function
	// We are passing the function itself to HandleFunc. That function is stored to be called later
	// when a matching request path is received
	// In a programming language with first-class functions, functions can be
	// assigned to variables, and then called from those variables
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/world", worldHandler)
	http.HandleFunc("/people", peopleHandler)
	// When you’re ready to open
	//apps up to requests from other computers, you can use a string of "0.0.0.0:8080" instead
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
