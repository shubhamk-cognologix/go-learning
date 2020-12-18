package main

import "fmt"

var flag = true

func main() {
	name := "Shubham"
	age := 24
	isMarried := false

	safe, speed, temp := true, 50, 32.65
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(flag)
	fmt.Println(safe)
	fmt.Println(speed)
	fmt.Println(temp)

	safe = false
}
