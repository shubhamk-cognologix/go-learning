package main

import "fmt"

type book struct {
	title string
	price int
}

type game struct {
	title string
	price int
}

func (b book) printBook() {
	fmt.Println(b)
}

func (g game) printGame() {
	fmt.Println(g)
}

func main() {

	var (
		b = book{"The conjuring", 560}

		g = game{"GTA", 680}
	)
	b.printBook()
	g.printGame()

	book.printBook(b)
	game.printGame(g)

}
