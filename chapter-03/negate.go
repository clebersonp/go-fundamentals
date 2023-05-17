package main

import "fmt"

func negate(value *bool) {
	*value = !*value
}

func main() {
	truth := true
	negate(&truth)
	fmt.Printf("Truth: %v\n", truth)

	lies := false
	negate(&lies)
	fmt.Printf("Lies: %v\n", lies)
}
