package main

import "fmt"

func main() {
	ints := [11]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, n := range ints {
		if n%2 == 0 {
			fmt.Printf("%d is even\n", n)
		} else {
			fmt.Printf("%d is odd\n", n)
		}
	}
}
