package main

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"

	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(loadedDeck))
	}

	os.Remove(filename)
}

func TestShuffle(t *testing.T) {
	d := newDeck()
	d.shuffle()

	if d[0] == "Ace of Spades" && d[1] == "Two of Spades" {
		t.Errorf("Your deck may not be shuffled correctly. Please run this test again. If problem persists, check the shuffle() fn.")
	}

	hand, _ := deal(d, 5)

	if len(hand) != 5 {
		t.Errorf("Expected hand length of 5, got %d", len(hand))
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()

	// Test with random length of hand each time
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	handLength := r.Intn(len(d) - 1)
	remLength := len(d) - handLength

	hand, remaining := deal(d, handLength)

	if len(hand) != handLength {
		t.Errorf("Expected hand length of %v, got %v", handLength, len(hand))
	}

	if len(remaining) != remLength {
		t.Errorf("Expected remaining length of %d, got %d", remLength, len(remaining))
	}
}
