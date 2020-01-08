package main

import "fmt"

// Interfaces are a way to create generic logic for different types
type bot interface {
	getGreeting() string // Every bot will have the getGreeting() fn
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

/* In a real example, we would put the greeting value
on the struct and make them the same type.
This is a contrived example. */
func (b englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string { // Can omit the variable for the receiver if it is not used
	return "Hola"
}

// Even though the bots are different types, they both have a getGreeting() fn
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
