package go_testing

import (
	"fmt"
	"testing"
)

// limiting the run test execution for some matches names of functions: go test -v -run Two
// go test -v -run Numbers
// go test -v -run Sum

func TestShouldSumTwoNumbers(t *testing.T) {
	num1, num2 := 10, 15
	sumExpected := num1 + num2
	result := Add(num1, num2)
	if sumExpected != result {
		t.Errorf(errorString(sumExpected, result))
	}
}

func TestShouldSumThreeNumbers(t *testing.T) {
	num1, num2, num3 := 10, 15, 10
	sumExpected := num1 + num2 + num3
	result := Add(num1, num2, num3)
	if sumExpected != result {
		t.Errorf(errorString(sumExpected, result))
	}
}

func errorString(expected int, current int) string {
	return fmt.Sprintf("expected '%d' but current '%d'", expected, current)
}
