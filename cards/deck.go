package main

import "fmt"

// Create a new type 'deck'
// which is a slice of strings

type deck []string

func (d deck) print() {
	for i, card := range d { // for index value, card in range cards
		fmt.Println(i, card)	// print index number and card
	}
}

func (d deck) printhello() {
	for i, card := range d { // for index value, card in range cards
		fmt.Println(i, card, "hello")	// print index number and card
	}
}