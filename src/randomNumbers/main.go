package main

import (
	"fmt"
	"math/rand"
)

func main1() {

	rand.Seed(10)
	guess := 10

	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
