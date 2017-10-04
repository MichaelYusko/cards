package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Test length of the array of cards
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// Test the first card of an array with cards
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of 'Ace of Spades' but got %v", d[0])
	}

	// Test the last card of an array with cards
	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected first card of 'Four of Clubs' but got %v", d[len(d) - 1])
	}
}


func TestSaveToDeckFromFile(t *testing.T) {

	// Remove the file if exists
	os.Remove("_decktesting")

	deck := newDeck()

	// Save into file
	deck.saveToFile("_decktesting")

	// Crete new deck from the file
	loadDeck := newDeckFromFile("_decktesting")

	// Check if length of the array greater than 16
	if len(loadDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadDeck))
	}

	// Remove the file after testing
	os.Remove("_decktesting")
}