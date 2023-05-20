package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// map literals
	// myApp := map[string]float64 {"a": 1.2, "b": 5.6}
	/*
		Multiline map literal:
		myApp := map[string]float64 {
			"a": 1.2,
			"b": 5.6, -> needs finished with comma
		}
	*/

	// As with slices, the zero value for the map variable itself is nil. If you
	// declare a map variable, but don’t assign it a value, its value will be nil. That
	// means no map exists to add new keys and values to. If you try, you’ll get a
	// panic: use make, it's more safety
	// Ex: var nilMap map[int]string
	// nilMap[3] = "three"
	// panic: assignment to entry in nil map
	// user make or map literal to create a map correctly

	// Removing key/value pair with delete
	// delete(map, key)

	filePath := os.Args[1]
	lines, err := getStrings(filePath)
	if err != nil {
		log.Fatal(err)
	}
	ranks := make(map[string]int)
	// The for...range loop handles maps in random order!
	// The for...range loop processes map keys and values in a random order
	// because a map is an unordered collection of keys and values. When you use
	// a for...range loop with a map, you never know what order you’ll get the  map’s contents in!
	for _, line := range lines {
		ranks[line]++
		//_, ok := ranks[line]
		//if !ok {
		//	ranks[line] = 1
		//} else {
		//	ranks[line]++
		//}
	}
	//fmt.Println(ranks)
	// print map using for range
	for key, value := range ranks {
		fmt.Printf("%v: %v\n", key, value)
	}
}

// GetStrings retrieves all lines from a file as a slice type
func getStrings(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}
