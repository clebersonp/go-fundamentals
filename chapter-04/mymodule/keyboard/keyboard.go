package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)

	// Multiple return values from a function or method
	// in Go, they can return any number of values. The most common use of multiple return values in Go is to return
	// an additional error value that can be consulted to find out if anything went wrong while
	// the function or method was running
	// Go does not allow us to declare a variable unless we use it
	// _ underscore(blank identifier) means we receive the return value, but we cannot use it, so compile will ignore that
	// short variable declaration: declare "input" and "err"
	input, err := reader.ReadString('\n')
	if err != nil { // condition, check if the return value for err is different of nil(null in other languages)
		return 0, err
	}

	input = strings.TrimSpace(input) // remove white spaces, tabs, break lines
	// short variable declaration: declare "grade" and assign "err" because it already exists above. Don't have issue for duplicated variable name for "err"
	grade, err := strconv.ParseFloat(input, 64) // convert string to float64
	if err != nil {
		return 0, err
	}
	return grade, nil
}
