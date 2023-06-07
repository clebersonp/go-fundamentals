package main

import (
	"fmt"
	"github.com/clebersonp/interfaces/with"
)

func main() {
	// when we implement the Stringer interface we don't need to format the message, just call any method that already treat it
	wheel := with.Wheel(4)
	fmt.Println(wheel)
}
