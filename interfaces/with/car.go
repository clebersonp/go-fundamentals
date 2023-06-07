package with

import "fmt"

type Car struct {
	Model string
}

func (c Car) Brake() {
	fmt.Printf("Braking a car '%v'\n", c.Model)
}
