package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func (d deck) saveToFile(filename string) error { // Deck receives and can return an error type
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // Convert string to byte slice. 0666 is read/write for anyone
}

func newDeckFromFile(filename string) deck {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err) // Log error
		os.Exit(1)                 // Exit the program
	}

	return deck(strings.Split(string(content), ",")) // Split type-converted byte slice on the string on the commas into a slice of strings
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // Create a random source from the nanosecond time
	r := rand.New(source)                           // Create a type Rand

	for i := range d {
		newPosition := r.Intn(len(d) - 1) // Get a new position integer in the deck from the random source

		d[i], d[newPosition] = d[newPosition], d[i] // Swap the two cards
	}
}
