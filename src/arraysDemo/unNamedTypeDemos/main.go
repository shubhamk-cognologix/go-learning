package main

import "fmt"

func main() {

	type bookcase [5]int

	blue := bookcase{3, 2, 3, 4, 5}
	red := [5]int{1, 2, 3, 4, 5}

	if blue == red {
		fmt.Printf("Equal.......!")
	} else {
		fmt.Printf("Not Equal.......!")
	}
}
