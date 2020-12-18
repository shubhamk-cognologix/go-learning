package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	moods := [2][3]string{
		{"happy", "good", "awesome"},
		{"sad", "bad", "terrible"},
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(3)

	switch os.Args[2] {
	case "positive":
		fmt.Printf("%s your mood is %s now", os.Args[1], moods[0][n])
	case "negative":
		fmt.Printf("%s your mood is %s now", os.Args[1], moods[1][n])
	default:
		fmt.Printf("%s your mood is %s now", os.Args[1], "neutral")
	}

}
