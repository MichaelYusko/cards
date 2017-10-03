package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
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

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Erorr: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")

	return deck(s)
}

// Shuffle cards
func (d deck) shuffle()  {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}