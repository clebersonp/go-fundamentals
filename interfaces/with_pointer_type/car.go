package with_pointer_type

import (
	"fmt"
	"strings"
)

type Car struct {
	Model string
}

func (c *Car) Brake() {
	c.Model = strings.ToUpper(c.Model)
	fmt.Printf("Braking a car '%v'\n", c.Model)
}

func (c *Car) GetModel() string {
	return c.Model
}

// CloseDoor is a concrete method
func (c *Car) CloseDoor() {
	fmt.Printf("Closing the '%v' doors\n", c.Model)
}
