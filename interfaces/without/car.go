package without

import "fmt"

type Car struct {
	Model string
}

func (c *Car) Accelerate() {
	fmt.Printf("Accelerating a car '%v'\n", c.Model)
}
