package main

import "fmt"

func main() {
	// Go’s runes make it easy to work with partial strings and not have to worry
	// about whether they contain Unicode characters or not. Just remember,
	// anytime you want to work with just part of a string, convert it to runes, not bytes
	// Golang works utf-8 for characters
	asciiString := "MARIA"
	utf8String := "лфцши"

	// each character has only 1 byte of space
	for index, currentByte := range asciiString {
		fmt.Printf("%d: %s\n", index, string(currentByte))
	}

	// each character has 2 bytes of space
	for index, currentByte := range utf8String {
		fmt.Printf("%d: %s\n", index, string(currentByte))
	}

	// len return the number of bytes
	// 1 byte == 8 bits
	// English character uses 1 byte for character
	// Russian uses 2 bytes for character
	fmt.Printf("English: %v, length: %d\n", asciiString, len(asciiString))
	fmt.Printf("Russian: %v, length: %d\n", utf8String, len(utf8String))

	// if convert a string to bytes and work with slice of it, it can be a nightmare
	// because the size of each language character can to go from 1 to 4 bytes
	// and sliced it the byte's slices can be lost character representation
	bytesAsciiString := []byte(asciiString)
	bytesUtf8String := []byte(utf8String)
	fmt.Printf("%v\n%v\n", bytesAsciiString, bytesUtf8String)
	for index, currentByte := range bytesAsciiString {
		fmt.Printf("%d: %s\n", index, string(currentByte))
	}
	for index, currentByte := range bytesUtf8String {
		fmt.Printf("%d: %s\n", index, string(currentByte))
	}

	sliceBytesAsciiString := bytesAsciiString[3:]
	sliceBytesUtf8String := bytesUtf8String[3:]
	fmt.Printf("%v\n%v\n", string(sliceBytesAsciiString), string(sliceBytesUtf8String))

	// the correct form is to use runes for safety characters conversion instead of bytes
	asciiRunes := []rune(asciiString)
	utf8Runes := []rune(utf8String)
	// now the size of slice will be the same because runes
	fmt.Printf("English: %v, Bytes: %v, length: %d\n", string(asciiRunes), asciiRunes, len(asciiRunes))
	fmt.Printf("Russian: %v, Bytes: %v, length: %d\n", string(utf8Runes), utf8Runes, len(utf8Runes))

	sliceAsciiRunes := asciiRunes[3:]
	sliceUtf8Runes := utf8Runes[3:]
	fmt.Printf("English: %v, Bytes: %v, length: %d\n", string(sliceAsciiRunes), sliceAsciiRunes, len(sliceAsciiRunes))
	fmt.Printf("Russian: %v, Bytes: %v, length: %d\n", string(sliceUtf8Runes), sliceUtf8Runes, len(sliceUtf8Runes))

	for index, currentByte := range asciiRunes {
		fmt.Printf("%d: %s\n", index, string(currentByte))
	}
	for index, currentByte := range utf8Runes {
		fmt.Printf("%d: %s\n", index, string(currentByte))
	}

}
