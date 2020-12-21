// Golang program to show how to
// declare and define the struct

package main

import "fmt"

// Address Defining a struct type
type Address struct {
	Name    string
	city    string
	Pincode int
}

func main() {

	// Declaring a variable of a `struct` type
	// All the struct fields are initialized
	// with their zero value
	var a Address
	fmt.Println(a)

	a1 := Address{"Akshay", "Dehradun", 3623572}

	fmt.Println("Address1: ", a1)

	a2 := Address{Name: "Anikaa", city: "Ballia",
		Pincode: 277001}

	fmt.Println("Address2: ", a2)

	a3 := Address{Name: "Delhi"}
	fmt.Println("Address3: ", a3)
}
