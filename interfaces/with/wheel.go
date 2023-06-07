package with

import (
	"fmt"
)

type Wheel int

// String is a Stringer interface method, and it's satisfy it
func (w Wheel) String() string {
	return fmt.Sprintf("The current number of wheels is %v", int(w))
}
