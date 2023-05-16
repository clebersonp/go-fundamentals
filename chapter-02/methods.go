package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// Go methods are kind of like the methods that you
	// may have seen attached to “objects” in other languages, but they’re a bit simpler
	var now time.Time = time.Now() // Now() is a function in the package time. Returns a Time type (structs like class in java)
	var year int = now.Year()      // Year() is a method, because we access it by the instance of Time with dot.
	// Methods are functions that are associated with values of a particular type (structs, obj in java)
	fmt.Println(year)

	// The dot indicates that the thing on its right belongs to the thing on its left.
	//Whereas the functions we saw earlier belonged to a package, the methods
	//belong to an individual value. That value is what appears to the left of the dot.
	broken := "G# r#cks!"
	fmt.Println("Broken:", broken)
	replacer := strings.NewReplacer("#", "o") // NewReplacer belongs to the package, it is a function
	fixed := replacer.Replace(broken)         // Replace is a method that belongs to the type(struct) Replacer
	fmt.Println("Fixed: ", fixed)

}
