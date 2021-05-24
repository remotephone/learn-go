package main

import (
	"os"
	"testing"
)

// This looks weird, but its the syntax for a test
func TestNewDeck(t *testing.T) {
	//Create new deck
	d := newDeck()
	//Check length
	if len(d) != 52 {
		// This is a formatted string, you need to give the %v so it knows where to include the string
		t.Errorf("The number of cards did not match, expected 52, got %v", len(d))
	}

	if d[0] != "Two of Spades" {
		t.Errorf("Wrong first card, expect Two Of Hearts but got %v", d[0])
	}
	
	if d[len(d) - 1] != "Ace of Clubs" {
		t.Errorf("Expected last card Ace of Clubs, got %v", d[len(d) - 1])
	}

}


func TestContainsFakeCard(t *testing.T) {
	d := newDeck()
	var testCard = "Hearts of S2even"
	var result = false
	for _, card := range d {
		if card == testCard {
			result = true
		}
	}
	if result == true {
		t.Errorf("Somethings wrong, fake card in deck")
	}
}

func TestContainsCard(t *testing.T) {
	d := newDeck()
	var testCard = "Seven of Hearts"
	var result = false
	for _, card := range d {
		if card == testCard {
			result = true
		}
	}
	if result == false {
		t.Errorf("This is good")
	}
}


func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("./_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}