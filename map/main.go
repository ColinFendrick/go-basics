package main

import "fmt"

func main() {
	colors := map[string]string{ // [key]value
		"red":   "#f00",
		"green": "#0f0",
		"white": "#fff",
	}

	colors["yellow"] = "nonsense"
	colors["blue"] = "#00f"
	delete(colors, "yellow")

	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Printf("Hex code for %v is %v \n", color, hex)
	}
}
