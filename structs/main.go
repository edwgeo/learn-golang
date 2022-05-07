package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// define a struct, which contains the properties and their types
type person struct {
	firstname string
	lastname  string
	contactInfo
}

func main() {
	jim := person{
		firstname: "Jim",
		lastname:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94600,
		},
	}

	// shortcut in Go - if our parameter is *person (as it is in receiver), Go will turn a person into a *person
	jim.updateName("jimmy")
	jim.print()
}

// *person is a '*' in front of a type ('person') - the '*' is part of a type description
// *pointerToPerson is a '*' in front of an actual pointer ('pointerToPerson') - the '*' is an operator
func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstname = newName // this doesn't work - why? because we are passing by value, so we only change a copy of the person that is passed in
	// you can make this work by putting *person in the receiver param, which says you want to receive the param as a pointer
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
