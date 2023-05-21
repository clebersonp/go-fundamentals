package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

// An interface is a set of methods that certain values are expected to have.
// You can think of an interface as a set of actions you need a type to be able to perform
// You define an interface type using the interface keyword, followed by
// curly braces containing a list of method names, along with any parameters or return values the methods are expected to have
// Any type that has all the methods listed in an interface definition is said to satisfy that interface.
// A type that satisfies an interface can be used anywhere that interface is called for
// The method names, parameter types (or lack thereof), and return value
// types (or lack thereof) all need to match those defined in the interface. A type can have methods in addition to those
// listed in the interface, but it must not be missing any, or it does not satisfy that interface
// A type can satisfy multiple interfaces, and an interface can (and usually should) have multiple types that satisfy it
// Many other languages would require us to explicitly say that MyType satisfies MyInterface.
// But in Go, this happens automatically. If a type has all the methods declared in an interface,
// then it can be used anywhere that interface is required, with no further declarations needed
// Interface types don’t describe what a value is: they don’t say what its
// underlying type is, or how its data is stored. They only describe what a value can do: what methods it has

type Player interface {
	Play(song string)
	Stop()
}

// If the number and types of all parameters and return values don’t match between a concrete type’s method definition
// and the method definition in the interface, then the concrete type does not satisfy the interface.

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}
func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}
