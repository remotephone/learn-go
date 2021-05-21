package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

// define a function with a receiver of deck type to return type string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// assign to bs and err the results of ioutil.ReadFile(filename_provided_as_argument)
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
// 	rand.Seed(time.Now().UnixNano())
// 	rand.Shuffle(len(d), func(i, j int) {d[i], d[j] = d[j], d[i]})
// }

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d{
		newPosition := r.Intn(len(d) - 1)
		d[i] , d[newPosition] = d[newPosition], d[i]
	}
}