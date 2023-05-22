package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func reportPanic() {
	p := recover() // if there's a panic it recovered and return it back
	// If recover is called when there is no panic, it simply returns nil
	if p == nil {
		return
	}
	// id there any value returned to p, then try to use type assertion and convert to an error that was passed to the panic function
	// but the type is not an error interface panic value can't be converted to an error type, and reportPanic won't print it.
	// example: panic("some message of error in string format")
	if err, ok := p.(error); ok {
		fmt.Println(err)
	} else {
		// if failed the conversion of type assertion of an error interface, so panic again with the same panic value without conversion
		panic(err)
	}
}

func scanDirectory(path string) {
	fmt.Println(path)
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			// recursive calls if the file is a directory
			scanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}

func main() {
	// call the reportPanic in defer mode(delay) to execute before the return this function but after all others calls
	defer reportPanic()

	// this is an issue because reportPanic won't be able to report it because it's a string into panic function instead of error interface
	// so, if the report won't be able to convert the panic value to an error interface, so it will panic again with the panic value
	// in this case will be a simple string
	//panic("Some other issue") // this call is just for test purpose

	// when scanDirectory encounters an error reading a directory, it
	// simply panics. All the recursive calls to scanDirectory exit
	scanDirectory("C:\\")
}
