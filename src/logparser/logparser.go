package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var lines int
	visitsMap := make(map[string]int)
	var domains []string
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		lines++
		fields := strings.Fields(in.Text())
		if len(fields) < 2 {
			fmt.Println("Wrong input")
			return
		}
		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Printf("Wrong count %q at line %d", fields[1], lines)
			return
		}
		if _, ok := visitsMap[domain]; !ok {
			domains = append(domains, domain)
		}
		visitsMap[domain] += visits
	}
	fmt.Printf("%-30s %-10s\n", "Domain", "Visits")
	fmt.Println(strings.Repeat("-", 45))
	for _, domain := range domains {
		fmt.Printf("%-30s %10d\n", domain, visitsMap[domain])
	}
}
