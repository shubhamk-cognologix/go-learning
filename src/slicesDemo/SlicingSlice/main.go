package main

import "fmt"

func main() {

	num := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	fmt.Println("\n\n Array: ", num[1:9])

	fmt.Println("\n\n Array > 5: ", num[6:])
	fmt.Println("\n\n Array < 5: ", num[:5])
}
