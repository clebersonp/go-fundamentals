package main

import (
	"fmt"
	"runtime/debug"
	"strings"
	"time"
)

func main() {

	channelWithoutPointer := make(chan string)
	channelWithPointer := make(chan *string)

	slicesSymbols := make([][]string, 0)
	groupOne := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	groupTwo := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	groupThree := []string{"M", "N", "O", "P", "R", "S", "T", "U", "W", "X"}

	slicesSymbols = append(slicesSymbols, groupOne)
	slicesSymbols = append(slicesSymbols, groupTwo)
	slicesSymbols = append(slicesSymbols, groupThree)

	fmt.Printf("Address slicesSymbols: %p, value: %v\n", &slicesSymbols, slicesSymbols)
	for _, groups := range slicesSymbols {
		fmt.Printf("Address group: %p, value: %v\n", &groups, groups)
	}

	fmt.Printf("===============Without Pointer=====================\n")
	for _, group := range slicesSymbols {
		GetGroupsWithoutPointer(group)
	}

	fmt.Printf("===============With Pointer=====================\n")
	for _, group := range slicesSymbols {
		GetGroupsWithPointer(&group)
	}

	fmt.Printf("===============Without Pointer And Channel=====================\n")
	for _, group := range slicesSymbols {
		go GetGroupsWithoutPointerAndChannel(channelWithoutPointer, group)
	}
	for i := 0; i < len(slicesSymbols); i++ {
		fmt.Println("==============================================")
		for j := 0; j < len(slicesSymbols[i]); j++ {
			fmt.Printf("%s,", <-channelWithoutPointer)
		}
		fmt.Printf("\n==============================================\n\n")
	}

	fmt.Printf("===============Slice Without Pointer And Channel With Pointer=====================\n")
	for _, group := range slicesSymbols {
		go GetGroupsAndChannelWithPointer(channelWithPointer, group)
	}
	for i := 0; i < len(slicesSymbols); i++ {
		fmt.Println("==============================================")
		for j := 0; j < len(slicesSymbols[i]); j++ {
			fmt.Printf("%s,", *<-channelWithPointer)
		}
		fmt.Printf("\n==============================================\n\n")
	}

	// this is not working.
	// to work it needs to pass the slice as a pass-copy-value, or create and store the slices into the root slices
	// already as pointer type
	fmt.Printf("===============Slice With Pointer And Channel With Pointer=====================\n")
	for _, group := range slicesSymbols {
		go GetGroupsWithPointerAndChannelWithPointer(channelWithPointer, &group)
	}
	for i := 0; i < len(slicesSymbols); i++ {
		fmt.Println("==============================================")
		for j := 0; j < len(slicesSymbols[i]); j++ {
			fmt.Printf("%s,", *<-channelWithPointer)
		}
		fmt.Printf("\n==============================================\n\n")
	}

	channelWithPointer2 := make(chan *string)

	slicesSymbolsWithPointer := make([]*[]string, 0)
	groupOneWithPointer := &[]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	groupTwoWithPointer := &[]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	groupThreeWithPointer := &[]string{"M", "N", "O", "P", "R", "S", "T", "U", "W", "X"}

	slicesSymbolsWithPointer = append(slicesSymbolsWithPointer, groupOneWithPointer)
	slicesSymbolsWithPointer = append(slicesSymbolsWithPointer, groupTwoWithPointer)
	slicesSymbolsWithPointer = append(slicesSymbolsWithPointer, groupThreeWithPointer)

	// this work fine.
	// The slices need to be created as pointer type
	fmt.Printf("===============All With Pointer=====================\n")
	for _, group := range slicesSymbolsWithPointer {
		go GetGroupsWithPointerAndChannelWithPointer(channelWithPointer2, group)
	}
	for i := 0; i < len(slicesSymbolsWithPointer); i++ {
		fmt.Println("==============================================")
		for j := 0; j < len(*slicesSymbolsWithPointer[i]); j++ {
			fmt.Printf("%s,", *<-channelWithPointer2)
		}
		fmt.Printf("\n==============================================\n\n")
	}

}

func GetGroupsWithoutPointer(group []string) {
	fmt.Println("==============================================")
	for _, element := range group {
		fmt.Printf("%s,", element)
		time.Sleep(5 * time.Millisecond)
	}
	fmt.Printf("\n==============================================\n\n")
}

func GetGroupsWithPointer(group *[]string) {
	fmt.Println("==============================================")
	for _, element := range *group {
		fmt.Printf("%s,", element)
		time.Sleep(5 * time.Millisecond)
	}
	fmt.Printf("\n==============================================\n\n")
}

func GetGroupsWithoutPointerAndChannel(channel chan string, group []string) {
	for _, element := range group {
		channel <- element
		time.Sleep(5 * time.Millisecond)
	}
}

func GetGroupsAndChannelWithPointer(channel chan *string, group []string) {
	fmt.Printf("Called by: %s, slices address: %p, value: %v\n", GetGoroutineId(), group, group)
	for _, element := range group {
		channel <- &element
		time.Sleep(5 * time.Millisecond)
	}
}

func GetGroupsWithPointerAndChannelWithPointer(channel chan *string, group *[]string) {
	fmt.Printf("Called by: %s, slices address: %p, value: %v\n", GetGoroutineId(), *group, group)
	for _, element := range *group {
		channel <- &element
		time.Sleep(5 * time.Millisecond)
	}
}

func GetGoroutineId() string {
	fields := strings.Fields(string(debug.Stack()))
	return fmt.Sprintf("%s - %s", fields[0], fields[1])
}
