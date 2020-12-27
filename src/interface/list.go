package main

import "fmt"

type printer interface {
	print()
}

type list []printer

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("List is empty")
		return
	}

	for _, it := range l {
		it.print()
	}
}

func (l list) discount(amount float64) {
	if len(l) == 0 {
		fmt.Println("List is empty")
		return
	}
	type discounter interface {
		discount(float64)
	}
	for _, it := range l {
		gm, ok := it.(discounter)
		if !ok {
			continue
		}
		gm.discount(amount)
	}
}
