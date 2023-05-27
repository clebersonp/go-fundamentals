package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {

	// ============================= Encoding ======================================
	m := Message{
		Name: "Alice",
		Body: "Hello",
		Time: 1294706395881547000,
	}

	// encoded json with marshal from encoding/json package
	b1, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	// b1 will be a []byte containing this JSON Data like b2:
	b2 := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
	stringOfB1 := string(b1)
	stringOfB2 := string(b2)
	fmt.Printf("Marshal result string b1: %v\nManual convertion string b2: %v\nB1 == B2? %t\n", stringOfB1, stringOfB2, stringOfB1 == stringOfB2)

	fmt.Println()
	// ============================= Decoding ======================================
	var m2 Message
	err = json.Unmarshal(b1, &m2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Unmarshal, decode JSON data to struct type: %#v\n\n", m2)

	var map1 map[string]any
	err = json.Unmarshal(b1, &map1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Convert JSON data to map: Name: %v, Body: %v, Time: %v\n\n", map1["Name"], map1["Body"], map1["Time"])

	// Unmarshal only decode field that fit in the structure, fields that not exists in the struct will be ignored.
	b3 := []byte(`{"Name": "Anna", "Food": "Chicken"}`)
	var m3 Message
	err = json.Unmarshal(b3, &m3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Unmarshal only knowed fields %#v\n\n", m3)

	// Generic JSON with interface{}
	var i interface{}
	i = "a string"
	i = 2011
	i = 2.777
	fmt.Printf("Type of var i: %T\n\n", i)

	// type assertion accesses the underlying concrete type:
	r, ok := i.(float64)
	if ok {
		fmt.Printf("Type of r: %T, value: %v\n\n", r, r)
	}

	// switch statements determines the type:

	switch v := i.(type) {
	case int:
		fmt.Println("V is an int: ", v)
	case float64:
		fmt.Println("V is a float: ", v)
	case string:
		fmt.Println("V is a string: ", v)
	default:
		fmt.Println("V is not of the types above")
	}
	fmt.Println()
	// ============================ Decoding arbitrary data =============================
	b4 := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err = json.Unmarshal(b4, &f)
	if err != nil {
		log.Fatal(err)
	}
	// type assertion to access f's underlying map[string]interface{}
	// In this way you can work with unknown JSON data while still enjoying the benefits of type safety
	m4 := f.(map[string]interface{})
	for k, v := range m4 {
		switch vv := v.(type) {
		case string:
			fmt.Println("Field", k, "is string", vv)
		case float64:
			fmt.Println("Field", k, "is float64", vv)
		case []interface{}:
			fmt.Println("Field", k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

}
