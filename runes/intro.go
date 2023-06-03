package main

import (
	"fmt"
	"reflect"
)

func main() {
	// rune is an alias for int32, and it is like char in java
	var letterA rune
	var letterB int32
	letterA = 'A'
	letterB = 'B'
	fmt.Printf("type of letterA: %T, type of letterB: %T, both types are the same? %t\n", letterA, letterB, reflect.TypeOf(letterA) == reflect.TypeOf(letterB))

	// emoji
	var emoji string
	emoji = "\U0001F64F"
	fmt.Println(emoji)
	fmt.Println("\U0001F389")
	fmt.Printf("%#v, %v, %v bytes\n", emoji, emoji, len(emoji))

	runesEmoji := []rune(emoji)
	runesEmoji = []rune{'\U0001F64F'}
	fmt.Printf("Raw value: %#v, Decimal: %v, unicode: %U\n", runesEmoji, runesEmoji, runesEmoji)

	emoji = string(runesEmoji) // conversion, typeWantToConvert(value)
	fmt.Println(emoji)

	fmt.Printf("Light tone: %v\n", "\U0001F3FB")
	fmt.Printf("Medium Light tone: %v\n", "\U0001F3FC")
	fmt.Printf("Medium tone: %v\n", "\U0001F3FD")
	fmt.Printf("Medium Dark tone: %v\n", "\U0001F3FE")
	fmt.Printf("Dark tone: %v\n", "\U0001F3FF")

	fmt.Printf("Thumb dark tone: %v\n", "\U0001F44D\U0001F3FF")

}
