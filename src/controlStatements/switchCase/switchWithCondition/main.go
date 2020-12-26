package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	switch {
	case now.Hour() < 12:
		fmt.Println("AM")

	default:
		fmt.Println("PM")
	}
}
