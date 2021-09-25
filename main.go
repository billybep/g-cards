package main

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()
	// cards := []string{ "Ace of Diamonds", newCard() }

	
	// hand, remainingCards := deal(cards, 5)
	
	// hand.print()
	// remainingCards.print()

	cards := newDeck()
	cards.saveToFile("my_cards")

}
