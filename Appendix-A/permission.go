package main

import (
	"fmt"
	"os"
)

func main() {
	// first argument, file name
	// second argument, int value for which operations will be doing to file in bitwise OR format
	// last one is the permission
	// documentation
	// go doc os OpenFile
	// go doc os O_CREATE
	// os.OpenFile("name_file.txt", os.O_CREATE|os.O_APPEND|os.O_CREATE, os.FileMode(0600))

	// os.FileMode permission works like in Unix and Mac system, but it will be ignored by Windows OS
	// permission for a class: user, group and other users
	// https://hpc-wiki.info/hpc/Chmod
	// types: '-' a file; 'd' a directory; '|' a link
	// type | owner | group | other
	// 0    |   7   |   0   |  0
	// -    |  rwx  |   -   |  -
	fmt.Println(os.FileMode(0700))
	fmt.Println(os.FileMode(0070))
	fmt.Println(os.FileMode(0007))
	// - | rwx | rwx | rwx
	fmt.Println(os.FileMode(0777))

	fmt.Println()

	fmt.Println("Octal notation")
	for i := 0; i <= 19; i++ {
		fmt.Printf("%3d:  %04o\n", i, i)
	}

	// any number preceded 0 will be threat as an octal number
	fmt.Printf("Decimal: %3d, Octal: %3o\n", 0777, 0777)

	// why unix permissions uses octal notation?
	// Because each digit of an octal number can be represented using just 3 bits of memory
	// Three bits is also the exact amount of data needed to store the permissions
	// for one user class (“user,” “group,” or “other”). Any combination of
	// permissions you need for a user class can be represented using one octal digit!
	// os.FileMode(0777)
	// 0 - type | 7 - user | 7 - group | 7 - other OCTAL NOTATION
	// 000      | 111      | 111       | 111  - BINARY NOTATION
	// -        | rwx      | rwx       | rwx  - EACH BIT IS IMPORTANT
	fmt.Printf("%09b\n", 0007)
	fmt.Printf("%09b\n", 0070)
	fmt.Printf("%09b\n", 0700)
	fmt.Printf("%09b\n", 0770)
	fmt.Printf("%09b\n", 0777)

	fmt.Println("Binary and permission representation")
	fmt.Printf("%09b %s\n", 0000, os.FileMode(0000))
	fmt.Printf("%09b %s\n", 0111, os.FileMode(0111))
	fmt.Printf("%09b %s\n", 0222, os.FileMode(0222))
	fmt.Printf("%09b %s\n", 0333, os.FileMode(0333))
	fmt.Printf("%09b %s\n", 0444, os.FileMode(0444))
	fmt.Printf("%09b %s\n", 0555, os.FileMode(0555))
	fmt.Printf("%09b %s\n", 0666, os.FileMode(0666))
	fmt.Printf("%09b %s\n", 0777, os.FileMode(0777))
}
