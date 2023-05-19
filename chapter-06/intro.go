package main

import (
	"fmt"
	"math"
)

func main() {

	// slice syntax: var mySlice []type . In slice type we don't specify the size
	// var mySlice []string // just declaration
	notes := make([]string, 7) // function make create a slice of 7 initial spaces
	fmt.Printf("Type: %T, len: %d, elements: %#v\n", notes, len(notes), notes)

	// slice literals
	// []type{value1, value2, ....}
	//[]int {9, 10, 12} alternative instead of call make function

	// Slice operator
	// mySlice := myArray[1:3] => 1:3 will be the slice operator and goes from index 1 until 3 exclusive(so 3-1=2)
	myArray := [5]string{"a", "b", "c", "d", "e"}
	mySlice := myArray[1:3]
	for _, letter := range mySlice {
		fmt.Println(letter)
	}

	// Slice has default for start and stop index
	//[:3] => starts from 0
	//[1:] => from index 1 until the last element

	// Change the underlying array, change the slice and vice-versa
	array1 := [5]string{"a", "b", "c", "d", "e"}
	slice := array1[0:3]
	fmt.Printf("array1: %#v\n", array1)
	fmt.Printf("slice: %#v\n", slice)
	array1[1] = "X"
	fmt.Printf("array1: %#v\n", array1)
	fmt.Printf("slice: %#v\n", slice)
	slice[2] = "W"
	fmt.Printf("array1: %#v\n", array1)
	fmt.Printf("slice: %#v\n", slice)

	// Because of these potential issues, you may find it’s generally better to
	// create slices using make or a slice literal, rather than creating an array and
	// using a slice operator on it. With make and with slice literals, you never
	// have to work with the underlying array

	// Builtin append() function append new elements to a slice and return a new one
	// For convention the return must be assigned to the same slice passed as append argument to avoid issues
	// Because there’s no easy way to tell whether the slice returned from append has the same underlying array as
	// the slice you passed in, or a different underlying array. If you keep both slices,
	// this can lead to some unpredictable behavior

	// Example without following the convention for assign return slices to the same variable
	s1 := []string{"s1", "s1"}
	s2 := append(s1, "s2", "s2")
	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")
	fmt.Println(s1, s2, s3, s4)
	s4[0] = "XX"
	fmt.Println(s1, s2, s3, s4)
	// Result:
	// [s1 s2] [s1 s2 s2 s2] [s1 s2 s2 s2 s3 s3] [s1 s2 s2 s2 s3 s3 s4 s4]
	// [s1 s2] [s1 s2 s2 s2] [XX s2 s2 s2 s3 s3] [XX s2 s2 s2 s3 s3 s4 s4]
	// s1 and s2 slices have a different underlying array, so the change can't be seen here!
	// s3 slice share an underlying array with s4

	// for convention just assign the return value back to the same slice variable you passed to append function.
	s5 := []string{"s5", "s5"}
	s5 = append(s5, "s6", "s6")
	s5 = append(s5, "s7", "s7")
	s5 = append(s5, "s8", "s8")
	fmt.Println(s5)
	s5[0] = "YY"
	fmt.Println(s5)

	// Slices and zero values
	// As with arrays, if your access a slice element that no value has been assigned to, you'll get the zero value
	// for that type back: like 0 Zero for numbers, false for booleans, nil for objs and "" empty for string
	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)
	fmt.Println(floatSlice[9], boolSlice[5])

	// Unlike arrays, the slice variable has a zero value: nil
	var intSlice []int
	var stringSlice []string
	fmt.Printf("intSlice: %#v, stringSlice: %#v\n, intSlice is nil? %t, stringSlice is nil? %t\n", intSlice, stringSlice, intSlice == nil, stringSlice == nil)

	// In other languages, that might require testing whether a variable actually contains a slice before attempting
	// to use it. But in Go, functions are intentionally written to treat a nil slice value as if it were an empty slice.
	// For example, the len function will return 0 if it’s passed a nil slice
	fmt.Println(len(intSlice))
	intSlice = append(intSlice, 27)
	fmt.Printf("intSlice: %#v\n", intSlice)

	// This means you generally don’t have to worry about whether you have an empty slice or a nil slice.
	// You can treat them both the same, and your code will “just work”!

	// Variadic functions
	// The last parameter of a variadic function receives the variadic arguments as
	// a slice, which the function can then process like any other slice
	// Only the last parameter in a function definition can be variadic; you
	// can’t place it in front of required parameters, same in java
	// func myFunc(param1 int, param2 ...string) {}
	// or:  func myFunc(param2 ...string) {}
	max1 := maximum(71.8, 56.2, 80.7)
	fmt.Printf("Max1 value: %v\n", max1)

	fmt.Println(inRange(1, 100, -12.5, 3.2, 0, 50, 103.5))

	fmt.Printf("Average: %.2f\n", average(1, 50.8, 15.25, 13, 0.75))

}

// maximum is a variadic function that check and return the max value in the given slice of numbers
func maximum(numbers ...float64) float64 {
	max := math.Inf(-1) // start with a very low value.
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

func average(numbers ...float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}
