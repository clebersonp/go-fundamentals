package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile(path string) (*os.File, error) {
	fmt.Println("Opening", path)
	return os.Open(path)
}
func CloseFile(file *os.File) {
	fmt.Println("Closing file", file.Name())
	_ = file.Close() // ignore explicit
}
func GetFloats(path string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(path)
	if err != nil {
		return nil, err
	}

	// defer is close to finally in java
	// defer works for function and method calls, only
	// Using defer ensures CloseFile will be called when GetFloats exits,
	// whether it completes normally or there’s an error parsing the file
	// Even if this function is given a file with bad data, it will still close the file
	//before exiting
	defer CloseFile(file) // defer(delay) making the function call until after the current function(GetFloats) exits.

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}

func main() {
	numbers, err := GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Printf("Sum: %.2f\n", sum)
}
