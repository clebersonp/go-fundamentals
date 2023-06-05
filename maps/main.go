package main

import (
	"fmt"
	"sort"
)

// existsInMap check if a given key exists in a given map with generic type
func existsInMap[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

func main() {

	// declare and create a map with short declaration and builtin make function
	ranks := make(map[string]int)
	fmt.Printf("Type: %T, Golang value: %#v, value: %v, is nil? %t\n", ranks, ranks, ranks, ranks == nil)

	// initialize map with literal
	ranks2 := map[string]int{"John": 2, "Maria": 15}
	fmt.Printf("Type: %T, Golang value: %#v, value: %v, is nil? %t\n", ranks2, ranks2, ranks2, ranks2 == nil)

	// Get a value of map with the key name
	// map key type can be almost any type, since we can compare with == operator
	// Zero-Value when access a key that hasn't been assigned to
	fmt.Println(ranks2["john"])

	// access the value by the existing key
	fmt.Println(ranks2["John"])

	// Zero-value for map variable is nil, and try access it will be to cause panic error
	var flags map[bool]int
	fmt.Printf("Type: %T, Golang value: %#v, value: %v, is nil? %t\n", flags, flags, flags, flags == nil)
	// panic: assignment to entry in nil map
	// flags[true] = 20

	// how to differentiate from Go's Zero-Value to not exists value for a specific key??
	// Golang can return a second value, it's a boolean value to tell us either the key exists or not
	counters := map[string]int{"Top": 10, "Bottom": 25, "Right": 0}
	leftCount := counters["Left"]
	fmt.Printf("Left count value: %v, Is it a Zero-Value or the key's value itself?? I don't know\n", leftCount)
	rightCount, ok := counters["Right"]
	fmt.Printf("Right count value: %v, Is it a Zero-Value or the key's value itself?? Now i know by the bool value. 'Right' key exists? %t\n", rightCount, ok)

	// we can just check if a specific key exists, without get its value
	_, ok = counters["Left"]
	fmt.Printf("'Left' key exists? %t\n", ok)

	// delete key/value pairs in the map with delete builtin function passing the map and the key
	evenNumbers := map[int]bool{2: true, 3: false, 4: true}
	delete(evenNumbers, 3)
	// nothing happens if a key does not exist
	delete(evenNumbers, 9)
	fmt.Printf("After remove from evenNumber map: %#v\n", evenNumbers)

	// Using range for loops with maps, return the key and value
	// The for...range lop processes map keys and values in a random order
	// because a map is an unordered collection of keys and values.
	scores := map[string]int{"John": 50, "Maria": 30, "Andrew": 25, "Anna": 12}
	for k, v := range scores {
		fmt.Printf("Key: %v, value: %v\n", k, v)
	}
	// to sort we need to extract keys values, sort them and iterate over map with the ordered keys again
	fmt.Println()
	var names []string
	for k, _ := range scores {
		names = append(names, k)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("Ordered by key. Key: %v, value: %v\n", name, scores[name])
	}

	fmt.Printf("Map scores, John key exists? %t\n", existsInMap(scores, "John"))
	fmt.Printf("Map scores, john key exists? %t\n", existsInMap(scores, "john"))
	fmt.Printf("Map even, 2 key exists? %t\n", existsInMap(evenNumbers, 2))
	fmt.Printf("Map counters, Top key exists? %t\n", existsInMap(counters, "Top"))

}
