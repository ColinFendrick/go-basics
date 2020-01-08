package main

import (
	"fmt"
	"math"
)

type sides map[string]float64 // This may be renamed to "edges". Personally I hate this but ymmv.

type polygon struct {
	name    string
	edges   sides
	areaFn  func(sides) float64
	perimFn func(sides) float64
}

type shape interface {
	area() float64
	perim() float64
	resize(n float64) polygon
	getName() string
}

func main() {
	circle := polygon{
		name:  "Circle",
		edges: sides{"radius": 4.0},
		areaFn: func(s sides) float64 {
			return math.Pi * math.Pow(s["radius"], 2)
		},
		perimFn: func(s sides) float64 { // Circle perimeter is calculated differently
			return math.Pi * 2 * s["radius"]
		},
	}

	rect := polygon{
		name:  "Rectangle",
		edges: sides{"s1": 2.0, "s2": 5.5},
		areaFn: func(s sides) float64 {
			return s["s1"] * s["s2"]
		},
	}

	triangle := polygon{
		name:  "Triangle",
		edges: sides{"s1": 3, "base": 4, "s2": 5},
		areaFn: func(s sides) float64 {
			return 0.5 * 3 * (s["base"])
		},
	}

	hexagon := polygon{
		name:  "Regular Hexagon",
		edges: sides{"s1": 4, "s2": 4, "s3": 4, "s4": 4, "s5": 4, "s6": 4},
		areaFn: func(s sides) float64 {
			return (3 * math.Sqrt(3) / 2) * math.Pow(s["s1"], 2)
		},
	}

	measure(circle, 4)
	measure(rect, 3)
	measure(triangle, 5)
	measure(hexagon, 2)
}

func (p polygon) perim() float64 {
	// If we need to create a custom perimeter function (ie for circles)
	if p.perimFn != nil {
		return p.perimFn(p.edges)
	}

	total := 0.0

	for _, s := range p.edges {
		total += s
	}

	return total
}

func (p polygon) area() float64 {
	return p.areaFn(p.edges)
}

func (p polygon) resize(n float64) polygon {
	newP := polygon{
		name:    p.name,
		edges:   sides{},
		areaFn:  p.areaFn,
		perimFn: p.perimFn,
	}

	for key, val := range p.edges {
		newP.edges[key] = val * n
	}

	return newP
}

func (p polygon) getName() string {
	return p.name
}

func measure(s shape, n float64) {
	name := s.getName()
	fmt.Println(name)
	fmt.Println("Area:", s.area())
	fmt.Println("Perim:", s.perim())

	newS := s.resize(n)
	fmt.Printf("\n%v Resized by %v:\n", name, n)
	fmt.Println("Area:", newS.area())
	fmt.Println("Perim:", newS.perim(), "\n ")
}
