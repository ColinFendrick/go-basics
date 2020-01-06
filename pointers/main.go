package main

import "fmt"

func main() {
	/*
		Slices create a position in memory that points to underlying array
		Because of this, when a function is called with a slice, we copy the slice data structure,
		which still has the reference to the original underlying array. So we can
		modify the underlying array, since the copy used inside the function (pass by value)
		still points to that array.

		This functionality is used in all non-primitive/reference types data types:
		slices, maps, channels, pointers, functions

		For primitives/value types, we must use pointers to reference the original, rather than the pass by value copy:
		int, float, string, bool, structs, arrays
	*/

	mySlice := []string{"Hi", "there", "how", "are", "you"}
	myArray := [2]int{0, 1}
	myString := "original string"

	// Will update the value
	updateSlice(mySlice)

	// Must use pointer in function to get these to work
	updateArray(&myArray)
	updateString(&myString)

	fmt.Printf("%v \n", mySlice)
	fmt.Printf("%v \n", myArray)
	fmt.Printf("%v \n", myString)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

func updateArray(a *[2]int) {
	a[0] = 1000
}

func updateString(s *string) {
	(*s) = "updated string"
}
