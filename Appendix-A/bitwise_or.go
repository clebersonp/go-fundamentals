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
	//os.OpenFile("name_file.txt", os.O_CREATE|os.O_APPEND|os.O_CREATE, os.FileMode(0600))

	// bitwise OR
	fmt.Printf("Int for RDONLY: \t%d\t%024b\n", os.O_RDONLY, os.O_RDONLY)
	fmt.Printf("Int for WRONLY: \t%d\t%024b\n", os.O_WRONLY, os.O_WRONLY)
	fmt.Printf("Int for   RDWR: \t%d\t%024b\n", os.O_RDWR, os.O_RDWR)
	fmt.Printf("Int for CREATE: \t%d\t%024b\n", os.O_CREATE, os.O_CREATE)
	fmt.Printf("Int for   EXCL: \t%d\t%024b\n", os.O_EXCL, os.O_EXCL)
	fmt.Printf("Int for  TRUNC: \t%d\t%024b\n", os.O_TRUNC, os.O_TRUNC)
	fmt.Printf("Int for APPEND: \t%d\t%024b\n", os.O_APPEND, os.O_APPEND)
	fmt.Printf("Int for   SYNC: \t%d\t%024b\n", os.O_SYNC, os.O_SYNC)
	fmt.Println()
	fmt.Println("Bitwise Or")
	fmt.Printf("\tO_WRONLY\t|\t%016b\t|\t%4d\n", os.O_WRONLY, os.O_WRONLY)
	fmt.Printf("\tO_CREATE\t|\t%016b\t|\t%4d\n", os.O_CREATE, os.O_CREATE)
	fmt.Printf("\tO_APPEND\t|\t%016b\t|\t%4d\n", os.O_APPEND, os.O_APPEND)
	fmt.Printf("\t--------\t|\t----------------\t|\t----\n")
	fmt.Printf("\t        \t \t%016b\t|\t%4d\n", os.O_WRONLY|os.O_CREATE|os.O_APPEND, int(os.O_WRONLY|os.O_CREATE|os.O_APPEND))
}
