package without

import "fmt"

type Motorcycle struct {
	Model string
}

func (m *Motorcycle) Accelerate() {
	fmt.Printf("Accelerating a motorcycle '%v'\n", m.Model)
}
