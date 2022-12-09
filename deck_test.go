package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Length is not 16, got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("First card is not Ace of Spades, got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Last card is not Four of Clubs, got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting!")

	deck := newDeck()
	deck.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 16 {
		t.Errorf("expected 16 cards in deck, got %v", len(loadedDeck))
	}
	os.Remove("_deckTesting")
}
