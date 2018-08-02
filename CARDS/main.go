package main

func main() {
	// cards := newDeck()
	// // []byte(cards.toString())
	// cards.saveToFile("myCards")

	// cards := newDeckFromFile("myCards")
	// cards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
