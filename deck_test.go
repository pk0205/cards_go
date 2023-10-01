package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	lenD := len(d)
	if lenD != 52 {
		t.Errorf("Expected deck length of 52 but got %d", lenD)
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades got %s", d[0])
	}

	if d[lenD-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs got %s", d[lenD-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deck-testing")
	d := newDeck()
	d.saveToFile("_deck-testing")

	loadedDeck := newDeckFromFile("_deck-testing")

	lenD := len(loadedDeck)
	if lenD != 52 {
		t.Errorf("Expected deck length of 52 but got %d", lenD)
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades got %s", d[0])
	}

	if d[lenD-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs got %s", d[lenD-1])
	}
	os.Remove("_deck-testing")
}
