package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newdeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got deck of length: %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	//delete file
	os.Remove("_decktesting")

	d := newdeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got deck of length: %d", len(d))
	}

	os.Remove("_decktesting")
}
