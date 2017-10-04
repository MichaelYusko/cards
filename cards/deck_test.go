package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Test length of the array of cards
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// Test first Cards
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of 'Ace of Spades' but got %v", d[0])
	}

	// Test last card
	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected first card of 'Four of Clubs' but got %v", d[len(d) - 1])
	}
}