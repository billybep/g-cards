package main

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()
	// cards := []string{ "Ace of Diamonds", newCard() }

	cards := deck{ "Ace of Diamonds", newCard() }
	cards = append(cards, "Six of Spades")

	cards.print()

	printState()
}

func newCard() string {
	return "Ace of Spades"
}
