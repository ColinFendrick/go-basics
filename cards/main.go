package main

import "fmt"

func main() {
	cards := newDeck()                    // Make a new deck
	hand, remainingDeck := deal(cards, 5) // Deal five cards
	hand.print()                          // Print the hand
	fmt.Println("\n----------- \nRemaining cards:\n ")
	remainingDeck.print()                     // Print the rest of the cards
	hand.saveToFile("hand")                   // Save the hand to file
	remainingDeck.saveToFile("remainingDeck") // Save the deck
	readHand := newDeckFromFile("hand")       // Get the hand from file
	fmt.Println("\nHand from file\n ")
	readHand.print() // Print the hand from file (same as hand above)

	newDeckFromFile("thisFileDoesNotExist") // To generate an error
}
