package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("data.txt") // open the data file for reading.
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file) // create a new scanner for the file
	var pounds []float64
	for scanner.Scan() { // read a line from the file
		text := scanner.Text()
		fmt.Printf("%#v\n", text) // print the line
		pound, err := strconv.ParseFloat(text, 64)
		if err != nil {
			log.Fatal(err)
		}
		pounds = append(pounds, pound)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	// cal the average
	size := float64(len(pounds))
	total := 0.0
	for _, pound := range pounds {
		total += pound
	}
	fmt.Printf("Total pounds: %.2f, beefs: %.2f, average: %.2f\n", total, size, total/size)
}
