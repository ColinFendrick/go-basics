package main

import "fmt"

func main() {
	cards := newDeck()                    // Make a new deck
	cards.shuffle()                       // Shuffle the cards
	hand, remainingDeck := deal(cards, 5) // Deal five cards

	handFileName, remainingFileName := "hand", "remainingDeck" // Create filenames
	hand.saveToFile(handFileName)                              // Save the hand to file
	remainingDeck.saveToFile(remainingFileName)                // Save the deck
	readHand := newDeckFromFile(handFileName)                  // Get the hand from file
	readRemaining := newDeckFromFile(remainingFileName)        // Read the remaining cards

	fmt.Println("\nHand:\n ")
	readHand.print() // Print the hand from file (same as hand above)
	fmt.Println("\n----------- \nRemaining cards:\n ")
	readRemaining.print()
}
