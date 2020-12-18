package main

import (
	"fmt"
	"os"
)

func main() {
	const (
		username = "shubham"
		password = "shubhs108"
	)

	if len(os.Args) < 3 {
		fmt.Println("Please provide credentials")
		return
	}
	if username != os.Args[1] {
		fmt.Println("Access denied for user ", os.Args[1])
	} else if password != os.Args[2] {
		fmt.Println("Invalid password !!!")
	} else {
		fmt.Println("Access granted")
	}
}
