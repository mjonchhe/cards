package main

func main() {
	// cards := newDeckFromFile("my_cards")
	cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// fmt.Println(cards.toString())
	cards.shuffle()
	cards.print()
	// cards.saveToFile("my_cards")
}
