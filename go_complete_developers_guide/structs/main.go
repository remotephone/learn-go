package main

import "fmt"

type contactInfo struct {
	email string
	zipCode	int
}

type person struct {
	firstName string
	lastName  string
}

type person2 struct {
	firstName string
	lastName  string
	contact contactInfo 
}


func main() {
	// first define all the properties it might have
	// create a ruleset and you can create a value
	// Let's create a person struct. First name, last name

	// Kinda wacky, relies on positional arguments, if you swap anything you break stuff
	alex := person{"Alex", "Anderson"}


	alex.print()
	// The ampersand is an operator. Look at this variable and look at this memory address its pointing to
	alexPointer := &alex
	alexPointer.updateName("alexa")
	alex.print()

}


func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// (p person) you can call this function on anything of type person
func (p person) print() {
	fmt.Printf("%+v", p)
}

