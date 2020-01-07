package main

import "fmt"

func main() {
	colors := map[string]string{ // [key]value
		"red":   "#f00",
		"green": "#0f0",
		"white": "#fff",
	}

	// var colors map[string]string <-- Can init zero-values with var
	// colors := make(map[string]string) <-- make() can also zero-value init a map

	colors["yellow"] = "nonsense" // Square brace syntax required on maps
	colors["blue"] = "#00f"
	delete(colors, "yellow") // delete() is a builtin golang fn

	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m { // Gets key, value pair in range of map
		fmt.Printf("Hex code for %v is %v \n", color, hex)
	}
}
