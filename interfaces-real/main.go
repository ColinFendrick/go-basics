package main

import "fmt"

type bot interface {
	greet() string
}

type languageBot struct{ language string }

func main() {
	e := languageBot{"English"}
	s := languageBot{"Spanish"}

	printGreeting(e)
	printGreeting(s)
}

func (l languageBot) greet() string {
	if l.language == "English" {
		return "Hello"
	} else if l.language == "Spanish" {
		return "Hola"
	} else {
		return "Hello"
	}
}

func printGreeting(b bot) {
	fmt.Println(b.greet())
}
