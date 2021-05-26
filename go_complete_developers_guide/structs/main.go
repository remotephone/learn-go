package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// first define all the properties it might have
	// create a ruleset and you can create a value
	// Let's create a person struct. First name, last name

	// Kinda wacky, relies on positional arguments, if you swap anything you break stuff
	alex := person{"Alex", "Anderson"}

	// more reasonable
	tim := person{firstName: "Tim", lastName: "Tomas"}

	// assigns zero value to each field
	var steve person

	fmt.Println(alex)
	fmt.Println(tim)
	fmt.Println(steve)
	fmt.Printf("%+v", steve)

	alex.firstName = "Alexson"
	fmt.Println(alex)

}
