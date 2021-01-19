package main

import "fmt"
import "io/ioutil"

// Create a new type 'deck'
// which is a slice of strings

type deck []string

// we did not add a receiver to this function because 
// we don't have a deck we're already working with
func newDeck() deck  {
	// create new variable of type deck that is empty
	// We could simply write out all 52 cards, but thats tedious
	cards := deck{}
	// Let's create two lists, one of suits, one of values
	// use two for loops to iterate through them and generate an entire deck!
	// for each suit in cardSuits
	//    for each value in cardValues
	//        Add a new card of 'value of suit' to 'cards' deck
	
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	// Iterate through both lists
	// The _ tells go we know theres a variable here but we're not gonna use it
	// Index, Value
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d { // for index value, card in range cards
		fmt.Println(i, card)	// print index number and card
	}
}

//declare function deal, d is of type deck, handSize of type int, returns two values, both of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}