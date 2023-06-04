package main

import "fmt"

func main() {
	// Unlike arrays, slices initialize with nil and need to instantiate before its use
	var names []string
	fmt.Printf("Type: %T, Size: %v, Golang value: %#v, value: %v\n", names, len(names), names, names)

	// try to use the only slices declaration(empty) will cause error
	// need to instantiate with make, Go builtin function or with append (will treat empty slice correctly)
	names = append(names, "John")
	fmt.Printf("Type: %T, Size: %v, Golang value: %#v, value: %v\n", names, len(names), names, names)

	// declare an empty slice of ages
	var ages []int
	// instantiate with make with initial capacity and ZERO-VALUES, Go's builtin function
	ages = make([]int, 3)
	fmt.Printf("Type: %T, Size: %v, Golang value: %#v, value: %v\n", ages, len(ages), ages, ages)

	// index is used to access slice element like array
	fmt.Printf("First element of names's slice: %v\n", names[0])

	// declaring and instantiate slice with short definition, and populate with builtin range function using for loop
	numbers := make([]int, 10)
	for index, _ := range numbers {
		numbers[index] += index * 2
	}
	fmt.Printf("Type: %T, Size: %v, Golang value: %#v, value: %v\n", numbers, len(numbers), numbers, numbers)

	// Instantiate slice with literals
	letters := []rune{'A', 'B', 'D', 'E'}
	fmt.Printf("Type: %T, Size: %v, Golang value: %#v, value: %v\n", letters, len(letters), letters, letters)

	// slice is built on top of an underlying array. The slice is merely a view into some of the array's elements
	// slice operator: anyArray[1:3] slice representation will be start from 1 until n - 1, 2 in this case
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	slice3 := underlyingArray[2:] // start from index 2 until the last one
	// trying slice operator with an incorrect index position will cause an out of bound for n-element array
	fmt.Printf("Type: %T, Size: %v, Golang value: %#v, value: %v\n", slice3, len(slice3), slice3, slice3)

	// Changing the underlying array, change the slice
	array1 := [5]string{"a", "b", "c", "d", "e"}
	slice1 := array1[0:3] // a, b, c
	slice2 := array1[2:]  // c, d, e
	fmt.Printf("Before: Type: %T, Size: %v, Golang value: %#v, value: %v\n", array1, len(array1), array1, array1)
	fmt.Printf("Before: Type: %T, Size: %v, Golang value: %#v, value: %v\n", slice1, len(slice1), slice1, slice1)
	fmt.Printf("Before: Type: %T, Size: %v, Golang value: %#v, value: %v\n", slice2, len(slice2), slice2, slice2)
	array1[2] = "X"
	fmt.Printf("After: Type: %T, Size: %v, Golang value: %#v, value: %v\n", array1, len(array1), array1, array1)
	fmt.Printf("After: Type: %T, Size: %v, Golang value: %#v, value: %v\n", slice1, len(slice1), slice1, slice1)
	fmt.Printf("After: Type: %T, Size: %v, Golang value: %#v, value: %v\n", slice2, len(slice2), slice2, slice2)

	// Because of these potential issues, you may find it’s generally better to create slices using make or a
	// slice literal, rather than creating an array and using a slice operator on it

	// slice type accept nil value, it's used to return function when any error can occur
	var nilSlice []string = nil
	fmt.Printf("Type: %T, Size: %v, Golang value: %#v, value: %v, It's nil? %t\n", nilSlice, len(nilSlice), nilSlice, nilSlice, nilSlice == nil)

}
