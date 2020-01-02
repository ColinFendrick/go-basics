package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of deck
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := [4]string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := [13]string{
		"Ace", "Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine", "Ten",
		"Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() { // Set deck type to receive print function, d is instance (like this or self), deck is type
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) { // Not using a receiver here
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string { // Deck receives this method
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
