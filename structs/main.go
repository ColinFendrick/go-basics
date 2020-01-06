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

	// & operator creates a pointer directly to the memory location
	jimPointer := &jim

	// The & is unnecessary here, we can just use jim directly
	jimPointer.updateFirstName("Jimmy")
	jim.updateFirstName("Jimmy")

	// Can also update the value directly if we want
	// Using either pointer or directly
	jimPointer.lastName = "UpdatedLastName"
	jim.lastName = "UpdatedLastName"

	jim.print()
	fmt.Println("\n ")
	alex.print()
}

// Go is a pass by value language, so when the function is called with the struct
// go copies that value and uses the copy for the function call, rather than the original.
// So to modify a value on the original we must use a pointer to reference the original value.
// *person means the function is taking a pointer to a person
func (p *person) updateFirstName(newName string) {
	// *p looks directly at the value of the pointer address
	(*p).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
