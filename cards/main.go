package main

func main() {
	cards := deck{newCard(), newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")

	cards.print()
	cards.printhello()
}

func newCard() string {
	return "Five of Diamonds"
}
