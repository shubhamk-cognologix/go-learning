package main

import "fmt"

func ptf(a *int) {
	*a = 748
}

func main() {

	var x = 100

	fmt.Printf("The value of x before function call is: %d\n", x)

	// var pa *int = &x
	// ptf(pa)
	ptf(&x)
	fmt.Printf("The value of x after function call is: %d\n", x)

}
