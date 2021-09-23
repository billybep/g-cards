package main

import "fmt"

// Create a new type of deck (slice of strings)
type deck []string

// Create function new deck
func newDeck() deck {
	cards := deck{}	// declare new empty variable

	cardSuits 	:= []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues	:= []string{"Ace", "One", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards

}

// Create function receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}