package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//func newDeck() deck {
//	cards := deck{"Ace of spades", "Two of Spades"}
//}