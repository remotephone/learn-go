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


func main1() {
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

	// Go will assign zero value to these struct values
	var jim person
	fmt.Printf("%+v", jim)
	alex.print()

	alexPointer := &alex
	alexPointer.updateName("alexa")
	alex.print()

}


func main2() {
	jim := person2{
		firstName: "Jim",
		lastName: "Jimson",
		contact: contactInfo{
			email: "jim@jim.com",
			zipCode: 94000,
		},
	}
	fmt.Printf("%+v", jim)

}


func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// (p person) you can call this function on anything of type person
func (p person) print() {
	fmt.Printf("%+v", p)
}

func main() {
	main1()
	main2()
}