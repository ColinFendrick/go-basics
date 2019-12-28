package main

import "fmt"

// Create a new type of deck
// which is a slice of strings
type deck []string

func (d deck) print() { // Set deck type to receive print function, d is instance (like this or self), deck is type
	for i, card := range d {
		fmt.Println(i, card)
	}
}
