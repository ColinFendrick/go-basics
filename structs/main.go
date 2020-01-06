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
	jim = person{"Jim", "Newname", contactInfo{"Jim@ghotmail.com", 33606}}

	// Can also define a struct by using var syntax
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Jones"
	alex.contactInfo = contactInfo{email: "Alex@infowars.com", zipCode: 12345}

	jim.print()
	fmt.Println("\n ")
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
