package with

import "fmt"

type Motorcycle struct {
	Model string
}

func (m Motorcycle) Brake() {
	fmt.Printf("Braking a motorcycle '%v'\n", m.Model)
}
