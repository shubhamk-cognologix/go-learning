package main

import (
	"fmt"
	"os"
)

func main() {

	switch os.Args[1] {
	case "Feb", "Mar", "Apr", "May":
		fmt.Println("This is summer season")
	case "Jun", "Jul", "Aug", "Sep":
		fmt.Println("This is Rainy season")
	case "Oct", "Nov", "Dec", "Jan":
		fmt.Println("This is winter season")
	default:
		fmt.Println("Invalid month, Please provide valid month")

	}
}
