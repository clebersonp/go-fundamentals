package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Naming a variable error would be a bad idea, because it would shadow the name of a type called error
	// shadowing names of builtin or not functions, packages, types etc. has precedence and will cause bad behaviour
	info, err := os.Stat("my.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(info.Size())
}
