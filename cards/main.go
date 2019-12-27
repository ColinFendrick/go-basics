package main

import "fmt"

func main() {
	cards := []string{newCard(), "Ace of Diamonds"} // Create a slice of strings
	cards = append(cards, "Six of spades")          // Append another element
	// Iterate over cards
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
