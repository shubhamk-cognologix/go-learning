package main

import "fmt"

func main() {
	var (
		minecrafts = game{product{title: "minecraft", price: money{20}}}
		tetris     = game{product{title: "tetris", price: money{10}}}
		mobydick   = book{product{title: "mobydick", price: money{5}}, "2011"}
		jigsaw     = puzzle{product{title: "jigsaw", price: money{8}}}
		doll       = toy{product{title: "doll", price: money{12}}}
	)

	var store list
	store = append(store, &minecrafts, &tetris, &mobydick, &jigsaw, &doll)

	store.print()

	// var p printer
	// p = mobydick
	// p.print()

	store.discount(2)
	fmt.Print("\n--------Discounted price--------")
	store.print()
	fmt.Println()
}
