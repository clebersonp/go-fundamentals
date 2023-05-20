package magazine

// Subscriber must be with first letter capitalized to export it to another packages
// For a type to be accessed outside the package it’s defined in, it must be exported: its name must begin with a capital letter
// Subscriber is an underlying struct type, and we can define type with it
type Subscriber struct {
	// Struct field names must also be capitalized if you want to export them from their package
	Name   string
	Rate   float64
	Active bool
	//HomeAddress Address // This is tedious, chaining type: s.HomeAddress.Street ... etc
	// An inner struct that is stored within an outer struct using an anonymous
	// field is said to be embedded within the outer struct. Fields for an embedded
	// struct are promoted to the outer struct, meaning you can access them as if
	// they belong to the outer struct
	// without the field name => anonymous field: s.Address.Street
	// Address struct type is embedded within the Subscriber
	// you don’t have to write out subscriber.Address.City to get at the City field; you can just write subscriber.City.
	// You can write the code as if there were no Address type at all;
	// it’s like the Address fields belong to the struct type they’re embedded within.
	Address
}

type Employee struct {
	Name   string
	Salary float64
	//HomeAddress Address
	Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
