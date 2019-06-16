package main

import (
	"os"
	"testing"
)

// newDeck tests:
// After we make a new deck, we need 52 items inside
// first card is Ace of Spades
// last card is King of Clubs

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got: %v", len(d))
	}

	if (d[0]) != "Ace of Spades" {
		t.Errorf("Expected card Ace of Spades at pos 0, but got: %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected deck length of 52, but got: %v", d[len(d)-1])
	}
}

func TestSaveToDeckandNewDeckFromFile(t *testing.T) {
	fileName := "_decktesting"
	os.Remove(fileName)
	deck := newDeck()

	deck.saveToFile(fileName)

	deck = newDeckFromFile(fileName)
	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got: %v", len(deck))
	}
	os.Remove(fileName)

}
