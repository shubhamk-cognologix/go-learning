package main

import "fmt"

func main() {

	// var m map[string]int
	m := map[string]int{
		"shubham": 86,
		"test1":   67,
		"atish":   80,
	}

	newMap := make(map[string]int)
	fmt.Printf("Shubham: %d", m["shubham"])

	for key, val := range m {
		fmt.Printf("\n Key: %s, Value: %d", key, val)
	}

	val, ok := newMap["Shubham"]

	fmt.Printf("\n Value: %d is present? %v", val, ok)

	fmt.Printf("\n Length of m:  %d", len(m))
	fmt.Printf("\n Length of newMap:  %d", len(newMap))

}
