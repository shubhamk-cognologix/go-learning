package main

import (
	"fmt"
	"strings"
)

func main() {
	const corpus = "lazy cat jumps again and again and again"
	words := strings.Fields(corpus)

	for i, v := range words {
		fmt.Printf("index: %d value: %s \n", i+1, v)
	}
}
