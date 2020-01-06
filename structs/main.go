package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "Jim@gmail.com",
			zipCode: 98101,
		},
	}

	// It is possible to define this by using the order of fields present in the struct definition
	jim = person{"Alex", "Jones", contactInfo{"Alex@gmail.com", 33606}}

	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
