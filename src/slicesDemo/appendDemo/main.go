package main

import "fmt"

func main() {
	var testSlice []int

	fmt.Printf("Type: %T", testSlice)
	fmt.Printf("\n Slice: %v", testSlice)

	testSlice = append(testSlice, 10)
	fmt.Printf("\n Slice:   %v", testSlice)
	fmt.Printf("\n Slice[0]:   %v", testSlice[0])

}
