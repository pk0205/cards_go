package main

func newCard() string {
	return "Queen of hearts"
}

func main() {
	cards := newDeckFromFile("my_cards")

	cards.shuffle()

	cards.print()

	// hand, remainingCards := cards.deal(5)

	// hand.print()
	// println("\n\n")
	// remainingCards.print()
}
