package main

import (
	"fmt"
	"math"
)

type square struct {
	sideLength float64
}

type triangle struct {
	height, base float64
}

type shape interface {
	getArea() float64
}

func main() {
	sq := square{
		sideLength: 4,
	}

	tri := triangle{
		base:   3,
		height: 4,
	}

	printArea(sq, "square")
	printArea(tri, "triangle")
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(sh shape, name string) {
	fmt.Printf("\nArea of %v: %v", name, sh.getArea())
}
