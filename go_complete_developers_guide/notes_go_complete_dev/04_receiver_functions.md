## Recievers

```go
func (d deck) print() {
```

Any  variable of type deck gets access to the print method

A `receiver` sets up methods of types we create.

```go

// SETUP CODE
func (d deck) print() {
	for i, card := range d { // for index value, card in range cards
		fmt.Println(i, card)	// print index number and card
	}
}


//Code from main.go
func main() {
	cards := deck{newCard(), newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")

	cards.print()
}
```

The word `deck` is a reference to the type we want to attach this to. The letter `d` is a "working" variable, or the  instance of the varaible we're workign with

When print is executed, cards is replaced with print function as d. 

By convention, the receiver is a 1 or 2 letter variation of the receiver. This is not a hrd requirement, but its a convention.
