package main

import (
	"fmt"
	"github.com/clebersonp/chapter-11/gadget"
)

// A variable declared with an interface type can hold any value whose type satisfies that interface
// gadget.TapePlayer and gadget.TapeRecorder satisfies all methods on gadget.Player
// Once you assign a value to a variable (or method parameter) with an
// interface type, you can only call methods that are specified by the interface on it

func playList(device gadget.Player, songs []string) {
	fmt.Printf("Executing the type of: %T\n", device)
	for _, song := range songs {
		// When you have a value of a concrete type assigned to a variable with an
		// interface type, a type assertion lets you get the concrete type back. It’s kind
		// of like a type conversion. Its syntax even looks like a cross between a
		// method call and a type conversion. After an interface value, you type a dot,
		// followed by a pair of parentheses with the concrete type. (Or rather, what
		// you’re asserting the value’s concrete type is
		// assertion fails will only occur at runtime not when compiling
		// When using type assertions, if you’re not absolutely sure which original
		// type is behind the interface value, then you should use the optional ok value
		// to handle cases where it’s a different type than you expected, and avoid a runtime panic
		if recorder, ok := device.(gadget.TapeRecorder); ok { // cast safety from an interface to concrete type to call more specialized method
			recorder.Record()
		} else {
			fmt.Println("Player was not a TapeRecorder. Record method will be skipped!")
		}
		device.Play(song)
	}
	device.Stop()
}

func main() {
	mixTape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	var player gadget.Player     // declare a variable using the interface type
	player = gadget.TapePlayer{} // values of TapePlayer satisfy Player, so we can assign this value to a variable with a type of Player
	playList(player, mixTape)

	// When Go decides whether a value satisfies an interface, pointer methods aren’t included for direct values.
	// But they are included for pointers. So the solution is to assign a pointer to a gadget.TapePlayer to the
	//Player variable, instead of a direct Switch value:
	// tape := gadget.TapePlayer{}
	//var p Player = & tape

	fmt.Println("======================")

	player = gadget.TapeRecorder{}
	playList(player, mixTape)
}
