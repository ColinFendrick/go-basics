package main

func main() {
	cards := deck{newCard(), "Ace of Diamonds"} // Create a slice of strings
	cards = append(cards, "Six of spades")      // Append another element

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
