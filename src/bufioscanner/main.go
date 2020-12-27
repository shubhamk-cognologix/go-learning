package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	lines := 0
	for in.Scan() {
		lines++
	}
	fmt.Println("No of Lines: ", lines)
}
