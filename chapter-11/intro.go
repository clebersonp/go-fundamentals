package main

import (
	"fmt"
	"log"
)

// ComedyError is a defined type with underlying type of 'string'
type ComedyError string

// Error satisfies the error interface, and we can assign it to a variable with the type error.
func (c ComedyError) Error() string {
	return string(c)
}

// OverheatError is a defined type with underlying type of 'float64'
type OverheatError float64

// Error satisfies the error interface, and we can assign it to a variable with the type error.
func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %.2f degress!\n", o)
}

// checkTemperature check whether actual temperature exceeded the safe temperature and return an error, otherwise return nil.
func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}

func main() {
	var err error = ComedyError("What's a programmer's favorite beer? Logger!")
	fmt.Println(err)

	err = checkTemperature(121.379, 100.0)
	if err != nil {
		log.Fatal(err)
	}
}
