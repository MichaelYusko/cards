package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

type deck []string

// Print all cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Create an array with cards
func newDeck() deck {
	cards := deck{}

	cardSuit := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuit {
		for _, value := range cardValue {
			cards = append(cards, value + " of " + suite)
		}
	}
	return cards
}

// Deal the cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Join by comma
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Save to file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}