package main

func main() {
	cards := newDeck()
	aHand, _ := deal(cards, 5)
	aHand.print()
	// cards.print()
}
