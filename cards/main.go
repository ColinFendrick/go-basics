package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	fmt.Println("\n----------- \nRemaining cards:\n ")
	remainingDeck.print()
	// cards.print()
	hand.saveToFile("hand")
	remainingDeck.saveToFile("remainingDeck")
}
