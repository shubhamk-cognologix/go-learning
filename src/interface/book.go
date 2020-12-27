package main

import (
	"fmt"
)

type book struct {
	product
	published interface{}
}

func (b *book) print() {
	b.product.print()
	fmt.Printf("\t Published: %v", b.published)
}
