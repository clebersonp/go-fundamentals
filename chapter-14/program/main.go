package main

import (
	"fmt"
	"github.com/clebersonp/chapter-14/prose"
)

func main() {
	fmt.Println(prose.JoinWithCommas([]string{"apple", "orange", "pear", "banana"}))
	fmt.Println(prose.JoinWithCommas([]string{"apple", "orange"}))
	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
}
