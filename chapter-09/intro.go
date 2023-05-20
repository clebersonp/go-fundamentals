package main

import (
	"fmt"
	"strings"
)

// Here’s a program that defines two new types, Liters and Gallons, both
// with an underlying type of float64. These are defined at the package level,
// so that they’re available within any function in the current package

// Liters is an underlying type of float64
// We can say: Liters is a type based on an underlying type of float64
type Liters float64

// Gallons is an underlying type of float64
// We can say: Gallons is a type based on an underlying type of float64
type Gallons float64

// Milliliters is an underlying type of float64
// We can say: Milliliters is a type based on an underlying type of float64
// 1/1000
type Milliliters float64

// MyType is an underlying type of string
type MyType string

// sayHi is a method for the MyType and just say "Hi" to the console
// The type of the receiver parameter is the type that the method becomes associated with
// But aside from that, the receiver parameter doesn’t get
// special treatment from Go. You can access its contents within the method
// block just like you would any other function parameter
// By convention, Go developers usually use a name consisting of a single
// letter—the first letter of the receiver’s type name, in lowercase. (This is
// why we used m as the name for our MyType receiver parameter
// Go uses receiver parameters instead of the “self”(python) or “this”(java) values seen in other languages
func (m MyType) sayHi() {
	fmt.Printf("%v is saying 'Hi'\n", m)
}

// CaseFormatter check if the value of MyType underlying type of string should to returned in the upper case or not
// methods is like a function, can receive arguments and return one or more values, the only difference is that
// it receive its type as a receiver, like de keyword 'this' or 'self' of others languages
func (m MyType) CaseFormatter(upperCase bool) string {
	if upperCase {
		return strings.ToUpper(string(m))
	}
	return strings.ToLower(string(m))
}

// ToGallons convert a type of Liters to Gallons both based on the same underlying type float64
// 1 liter == 0.264 gallon
//func ToGallons(l Liters) Gallons {
//	return Gallons(l * 0.264)
//}
// Defining a method with the same name of others types methods

// ToGallons convert a type of Liters to Gallons both based on the same underlying type float64
// 1 liter == 0.264 gallon
func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

// ToGallons convert a type of Milliliters to Gallons both based on the same underlying type float64
// 1 gallon == milliliters * (0.264/1_000)
// Experience with other languages told us that having a variety of methods with the same name but
// different signatures was occasionally useful but that it could also be
// confusing and fragile in practice. The Go language is simplified by not supporting overloading, and so it does not support it
// Go team made similar decisions in other areas of the language, too;
// when they have to choose between simplicity and adding more features, they generally choose simplicity
// Fixing the function name conflict using methods
//func ToGallons(m Milliliters) Gallons {
//	return Gallons(m * (0.264 / 1_000))
//}

// ToGallons convert a type of Milliliters to Gallons both based on the same underlying type float64
// 1 gallon == milliliters * (0.264/1_000)
// A method definition is very similar to a function definition. In fact, there’s
// really only one difference: you add one extra parameter, a receiver
// parameter, in parentheses before the function name.
// method syntax: func (referenceType) nameMethod() returnType if used {}
// The value you’re calling the method on is known as the method receiver.
// The similarity between method calls and method definitions can help you
// remember the syntax: the receiver is listed first when you’re calling a
// method, and the receiver parameter is listed first when you’re defining a method
// The name of the receiver parameter in the method definition isn’t important,
// but its type is; the method you’re defining becomes associated with all values of that type
func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * (0.264 / 1_000))
}

// ToLiters convert a type of Gallons to Liters both based on the same underlying type float64
// 1 gallon == 3.785 liters
// This is a simple function and not a method, because it's not associated for a receiver type
//func ToLiters(g Gallons) Liters {
//	return Liters(g * 3.785)
//}

// ToLiters convert a type of Gallons to Liters both based on the same underlying type float64
// 1 gallon == 3.785 liters
// This is a method for the defined type Gallons of underlying type float64
func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

// ToMilliliters convert type of Gallons to Milliliters both based on the same underlying type float64
func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * (3.785 * 1_000))
}

// Number is a type of underlying type of int
type Number int

// Double is a method of type Number of underlying type int and try to double its value
// but we need a pointer reference because Go 'pass-copy-value' for methods too
// And like any other parameter, a receiver parameter receives a
// copy of the receiver value. If you make changes to the receiver within a method, you’re changing the copy, not the original
// You can only define new methods on types that were declared in
// the current package. Defining a method on a globally defined type like int will result in a compile error
func (n *Number) Double() {
	*n *= 2 // needs to update the value at the pointer and not the pointer address itself
}

func main() {

	// Go defined types most often use structs as their underlying types, but
	// they can also be based on ints, strings, booleans, or any other type
	var carFuel Gallons
	var busFuel Liters

	// Once you’ve defined a type, you can do a conversion to that type from any
	// value of the underlying type. As with any other conversion, you write the
	// type you want to convert to, followed by the value you want to convert in parentheses
	carFuel = Gallons(10.0) // convert a float64 to Gallons underlying type
	busFuel = Liters(240.0) // convert a float64 to Liters underlying type
	fmt.Println(carFuel, busFuel)

	// use short variable declarations together with type conversion
	myCarCapacityInLiters := Liters(38.5)
	fmt.Println(myCarCapacityInLiters)

	// If you have a variable that uses a defined type, you cannot assign a value of
	// a different defined type to it directly, even if the other type has the same underlying type
	// But you can convert between types that have the same underlying type, in this case both are float64
	// But Go only considers the value of the underlying type when doing a conversion;
	// there is no difference between Gallons(Liters(240.0)) and Gallons(240.0).
	var l1 Liters
	var g1 Gallons
	l1 = Liters(Gallons(10.0) * 3.785) // convert from gallons to liters
	g1 = Gallons(Liters(35.0) * 0.264) // convert from liters to gallons
	fmt.Println(l1, g1)

	// A defined type supports all the same operations as its underlying type.
	// Types based on float64, for example, support arithmetic operators like +,
	// -, *, and /, as well as comparison operators like ==, >, and <.
	fmt.Println(Liters(1.2) + Liters(3.4))
	//fmt.Println(Liters(1.2) + Gallons(3.4)) // don't compile because the two types are different even they have the same underlying type float64
	fmt.Println(Gallons(9) / Gallons(3))
	fmt.Println(Gallons(5.2) == Gallons(3))
	fmt.Println(Gallons(5.2) >= Gallons(3))
	fmt.Println(Gallons(5.2) < Gallons(3))
	fmt.Println(Gallons(5.2) - Gallons(3))

	// A defined type can be used in operations together with literal values because the underlying type are the same
	fmt.Println(Liters(25.6) - 7.98)

	fmt.Printf("Gallons to liters: %.2f\n", Gallons(7.5).ToLiters())
	fmt.Printf("Liters to gallons: %.2f\n", Liters(29.98).ToGallons())           // Liters(29.98) is the receiver on the method
	fmt.Printf("Milliliters to gallons: %.2f\n", Milliliters(5_000).ToGallons()) // calling the conversion method on Milliliters type

	// we can call value as the receiver, receiver can call method of its type
	// Once a method is defined on a type, it can be called on any value of that type.
	value := MyType("a MyType value")
	value.sayHi()

	m1 := MyType("joHn addAN")
	result := m1.CaseFormatter(true)
	fmt.Println(result)

	// Using pointer to update its own value type inside a method
	n := Number(50)
	fmt.Println(n)

	// When you call a method that requires a pointer receiver on a variable with a nonpointer type,
	// Go will automatically convert the receiver to a pointer for you. The same is
	// true for variables with pointer types; if you call a method requiring a value
	// receiver, Go will automatically get the value at the pointer for you and pass that to the method
	n.Double()
	fmt.Println(n)

	// You can only get pointers to values that are stored in variables. If you
	// try to get the address of a value that’s not stored in a variable, you’ll get an error
	//&Number(10) // don't compile

	// The same limitation applies when calling methods with pointer
	// receivers. Go can automatically convert values to pointers for you, but
	// only if the receiver value is stored in a variable. If you try to call a
	// method on the value itself, Go won’t be able to get a pointer, and you’ll get a similar error
	// Instead, you’ll need to store the value in a variable, which will then allow Go to get a pointer to it
	//Number(10).Double() // don't compile

	milk := Gallons(2)
	fmt.Printf("%0.3f gallons equals %0.3f liters\n", milk, milk.ToLiters())
	fmt.Printf("%0.3f gallons equals %0.3f milliliters\n", milk, milk.ToMilliliters())

}
