package main

import "fmt"

func main() {
	a := 10

	var s *int
	fmt.Println("------------Before------------")
	fmt.Println("S: ", s)
	s = &a
	fmt.Println("------------After-------------")
	fmt.Println("A: ", a)
	fmt.Println("&A: ", &a)
	fmt.Println("S: ", s)
	fmt.Println("S: ", *s)

	*s = 20
	fmt.Println("------After value change------")
	fmt.Println("A: ", a)
	fmt.Println("&A: ", &a)
	fmt.Println("S: ", s)
	fmt.Println("S: ", *s)

}
