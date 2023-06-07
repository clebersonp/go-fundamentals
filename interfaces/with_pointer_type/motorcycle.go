package with_pointer_type

import (
	"fmt"
	"strings"
)

type Motorcycle struct {
	Model string
}

func (m *Motorcycle) Brake() {
	m.Model = strings.ToUpper(m.Model)
	fmt.Printf("Braking a motorcycle '%v'\n", m.Model)
}

func (m *Motorcycle) GetModel() string {
	return m.Model
}

// CloseHelmetVisor is a concrete method
func (m *Motorcycle) CloseHelmetVisor() {
	fmt.Printf("Closing the pilot's helmet visor from the '%v' motorcycle\n", m.Model)
}
