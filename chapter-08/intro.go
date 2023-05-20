package main

import "fmt"

func main() {
	// You can use a struct type as the type of a variable you’re declaring
	// For every new variable with same fields we have to declare all the same struct again and again...
	// It's more common to use a defined type to declare struct variables
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}

	// myOtherSubscriber variable that hold structs
	var mySubscriber struct {
		name   string
		rate   float64
		active bool
	}

	/*
		A struct (short for “structure”) is a value that is constructed out of other
		values of many different types. Whereas a slice might only be able to hold
		string values or a map might only be able to hold int values, you can
		create a struct that holds string values, int values, float64 values, bool
		values, and more—all in one convenient group.

		Struct is like a blueprint class in java

		struct {
			field1 string
			field2 int
		}
	*/

	fmt.Printf("Value: %#v, type: %T\n", myStruct, myStruct)

	// dot operator to indicate fields that "belong to" a struct.
	myStruct.number = 3.14
	myStruct.word = "Hello World!"
	myStruct.toggle = true
	fmt.Printf("%v\n", myStruct)

	mySubscriber.name = "Aman Singh"
	mySubscriber.rate = 4.99
	mySubscriber.active = true
	fmt.Println("Name:", mySubscriber.name)
	fmt.Println("Monthly rate:", mySubscriber.rate)
	fmt.Println("Active?", mySubscriber.active)

	fmt.Println("=================================")
	var porsche car // type car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println("Name:", porsche.name)
	fmt.Println("Top speed:", porsche.topSpeed)

	var bolts part = part{
		description: "first description",
		count:       1,
	}
	fmt.Printf("Part => Type: %T, Values: %v\n", bolts, bolts)
	bolts.description = "Hex bolts"
	bolts.count = 24
	fmt.Println("Description:", bolts.description)
	fmt.Println("Count:", bolts.count)
	showInfo(&bolts) // access the memory address to pointer
	order1 := miniOrder("My minimum order")
	fmt.Println(order1)
	fmt.Println("=================================")

	var subscriber1 myOtherSubscriber
	subscriber1.name = "Aman Singh"
	var subscriber2 myOtherSubscriber
	subscriber2.name = "Beth Ryan"
	fmt.Printf("Sub1 name: %v, Sub2 name: %v\n", subscriber1.name, subscriber2.name)

	fmt.Println("=================================")
	// Shadow the type name, don't use an existing type name as a variable name!
	/*var part part
	part.description = "Description"
	fmt.Println(part)
	var part2 part // refer to the variable name instead o type car, this is the shadow
	part2.description = "Another description"*/

	fmt.Println("=================================")
	// function with a pointer reference to update the type structure field, because Go pass-copy-value instead of reference or pointer
	subscriber3 := myOtherSubscriber{
		name:   "Example of pointer to change the value",
		rate:   1_000,
		active: true,
	}
	fmt.Printf("Before change: %v\n", subscriber3)
	applyDiscount(&subscriber3)
	fmt.Printf("After change: %v\n", subscriber3)

}

// type definition
// syntax: keyword type nameOfType keyword struct(underlying type)
// type myType struct { //fields... }
// Just like variables, type definitions can be written within a function. But
// that will limit its scope to that function’s block, meaning you won’t be able
// to use it outside that function. So types are usually defined outside of any
// functions, at the package level
type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

// subscribe type at the package level, it's better than at the function level and variable declaration
type myOtherSubscriber struct {
	name   string
	rate   float64
	active bool
}

// using pointers ensures that only one copy of each struct needs to be kept in
// memory, while still allowing the program to work as normal
func showInfo(p *part) {
	fmt.Println("Description:", p.description)
	fmt.Println("Count:", p.count)
}

// using pointers ensures that only one copy of each struct needs to be kept in
// memory, while still allowing the program to work as normal
func miniOrder(description string) *part {
	var p part
	p.description = description
	p.count = 100
	return &p
}

// Go is a “pass-by-value” language, meaning
// that function parameters receive a copy of the arguments the function was
// called with. If a function changes a parameter value, it’s changing the copy,
// not the original
// The same thing is true for structs. When we pass a type struct to
// some function, the function receives a copy of the struct. So when we set
// some field on the struct, we’re modifying the copied struct, not the
// original
// When calling the function, we used the address-of operator (&) to pass a pointer to
// the value we wanted to update. Then, within the function, we used the *
// operator to update the value at that pointer
// We can use pointers to allow a function to update a struct as well
// Only functions that modify existing structs without returning them have to use pointers for those changes to be preserved
// Unless your struct has only a couple small fields, it’s often a
// good idea to pass functions a pointer to a struct, rather than the struct itself.
// (This is true even if the function does not need to modify the struct.) When
// you pass a struct pointer, only one copy of the original struct exists in
// memory. The function just receives the memory address of that single
// struct, and can read the struct, modify it, or whatever else it needs to do, all without making an extra copy
// using pointers ensures that only one copy of each struct needs to be kept in
// memory, while still allowing the program to work as normal
func applyDiscount(s *myOtherSubscriber) {
	// The dot notation to access fields works on struct pointers as well as the structs themselves
	// That means we don't have to use the operator *(value-at operator) to access the pointer value
	// Or we can access by * operator, but need wrapper the variable reference with the operator * to access the field
	// And that’s how the applyDiscount function is able to update the struct
	// field without using the * operator. It assigns to the rate field through the struct pointer
	s.rate = 4.99 // this approach is better than that bellow, assign to the struct field through the pointer.
	(*s).rate = 5.01
}
