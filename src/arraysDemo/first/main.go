package main

import "fmt"

const (
	winter = 1
	summer = 4
	yearly = winter + summer
)

func main() {
	var books [yearly]string

	books[0] = "Shubham"
	books[1] = "Prateek"
	books[2] = books[0] + "Kumbhar"
	books[4] = books[1] + "walkhade"

	fmt.Printf("Type: %T ", books)
	fmt.Println("\n Books: ", books)
	fmt.Printf("\n Books: %q ", books)
	fmt.Printf("\n Books: %s ", books)
}
