package main

import "fmt"

func main() {
	const (
		max = 1000
	)
	var sum int
	for i := 0; i <= max; i++ {
		sum += i
	}
	fmt.Println("Sum: ", sum)

	//Sum of even numbers
	var i int
	sum = 0
	for {
		if i > max {
			break
		}
		if i%2 == 0 {
			sum += i
		}
		i++
	}
	fmt.Println("Sum of even numbers: ", sum)

	sum = 0
	i = 0
	for {
		if i > max {
			break
		}
		if i%2 == 0 {
			i++
			continue
		}
		sum += i
		i++
	}
	fmt.Println("Sum of Odd numbers: ", sum)
}
