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
	fmt.Println(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m { // Gets key, value pair in range of map
		fmt.Printf("Hex code for %v is %v \n", color, hex)
	}
}

/* Map vs struct:
1. Keys/values must be of same type in map. Values can be different in struct.
2. Maps can be iterated over. Keys are indexed. Cannot iterate over struct.
3. Map is a reference type. Struct is a value type. We can directly modify the map values in a func.
4. Map does not need to know list of all keys at compile time.
We can add initialize with zero-values and add values later.
Struct must be defined beforehand.

Map is usually for closely related properties. Struct might represent a bunch of different properties.
*/
