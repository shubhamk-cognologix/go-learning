package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	words := make(map[string]bool)

	for in.Scan() {
		word := strings.ToLower(in.Text())

		if len(word) > 2 {
			words[word] = true
		}
	}

	query := os.Args[1]

	if _, ok := words[query]; ok {
		fmt.Printf("Input contains the word %q \n", query)
		return
	}
}
