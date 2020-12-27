package main

import "fmt"

type money struct {
	amt float64
}

type product struct {
	title string
	price money
}

func (m *money) string() string {
	return fmt.Sprintf("$%.2f", m.amt)
}
func (p *product) print() {
	fmt.Printf("\n %-15s: %s", p.title, p.price.string())
}

func (p *product) discount(amount float64) {
	p.price.amt = p.price.amt - amount
}
