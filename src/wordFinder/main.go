package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again"

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]

queries:
	for _, val := range query {
	search:
		for i, word := range words {
			switch val {
			case "and", "or", "the":
				break search
			}
			if val == word {
				fmt.Printf("Word %s found at %d \n", val, i+1)
				break queries
			}
		}
	}
}
