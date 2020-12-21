package main

import "fmt"

func main() {
	var a int = 1000
	var aPtr *int = &a
	var dPtr **int = &aPtr

	fmt.Println("\n A: ", a)
	fmt.Println("\n A: ", &a)

	fmt.Println("\n --------Single pointer---------")
	fmt.Println("\n APtr: ", aPtr)
	fmt.Println("\n APtr: ", *aPtr)

	fmt.Println("\n --------Double pointer---------")
	fmt.Println("\n A: ", **dPtr)
}
