package main

import "fmt"

func main() {

	const (
		USD = iota
		INR
		EUR
		AED
	)
	forexRates := [...]float64{
		USD: 1.0,
		INR: 70.25,
		EUR: 1.15,
		AED: 5.75,
	}
	fmt.Println(INR, USD, EUR, AED)
	fmt.Printf("%g \n", forexRates)
	fmt.Printf("%g \n", forexRates[USD])

}
