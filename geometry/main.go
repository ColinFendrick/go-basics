package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
	name() string
}

type greeting func() string // Can define a function as a type if it is reused

type rect struct {
	width, height float64
	greet         greeting // Can define a method as a field
}

type circle struct {
	radius float64
	greet  greeting
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perimeter() float64 {
	return 2*r.height + 2*r.width
}

func (rect) name() string {
	return "Rectangle"
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return math.Pi * 2 * c.radius
}

func (circle) name() string { // Standard way to have a struct receive the function
	return "Circle"
}

/* We cannot use the greet() method inside of hjere, since this is a field
rather than a received metnod, and does not exist before the specific instances are created.
In a real world scenario, if we knew every shape would have a greet() method,
we would put it as a received fn above and put it in the interface.
*/
func measure(s shape) {
	fmt.Println(s.name())
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	r := rect{
		width:  3,
		height: 5,
		greet: func() string { // Can define a method as a field
			return "Hello, I am a rectangle"
		},
	}

	c := circle{
		radius: 4,
		greet: func() string {
			return "Hello, I am a circle"
		},
	}

	fmt.Println(r.greet())
	measure(r)
	fmt.Println(c.greet())
	measure(c)
}
