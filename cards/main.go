package main

import "fmt"

func main() {
	cards := newDeck()
	aHand, bHand := deal(cards, 5)
	aHand.print()
	fmt.Println("-----------")
	bHand.print()
	// cards.print()
}
