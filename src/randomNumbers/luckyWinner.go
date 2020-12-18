package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const maxTerns = 10
	rand.Seed(time.Now().UnixNano())
	guess := 10
	flag := false

	for turn := 0; turn <= maxTerns; turn++ {
		n := rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
		if n == guess {
			flag = true
			fmt.Println("\n YOU WIN.......")
			break
		}
	}
	if !flag {
		fmt.Println("YOU Lost.......")
	}
}
