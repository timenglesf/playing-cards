package playingcards

import (
	"reflect"
	"testing"

	"github.com/timenglesf/playing-cards/internal/assert"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	assert.Equal(t, len(deck), 52)
	for _, card := range deck {
		if card.Suit < Diamonds || card.Suit > Spades {
			t.Errorf("Invalid suit: %d", card.Suit)
		}
		if card.Rank < Ace || card.Rank > King {
			t.Errorf("Invalid rank: %d", card.Rank)
		}
	}
}

func TestShuffleDeck(t *testing.T) {
	deck := NewDeck()
	shuffled := NewDeck()
	shuffled.Shuffle()

	assert.Equal(t, len(deck), 52)
	assert.Equal(t, len(shuffled), 52)
	assert.Equal(t, len(deck), len(shuffled))

	matchCount := 52
	for i := range deck {
		if deck[i] == shuffled[i] {
			matchCount--
		}
	}

	if matchCount == 0 {
		t.Errorf("Deck not shuffled")
	}
}

func TestReshuffleDeck(t *testing.T) {
	deck := NewDeck()
	shuffled := NewDeck()
	shuffled.Shuffle()
	shuffled.Reshuffle()
	assert.Equal(t, len(deck), 52)
	assert.Equal(t, len(shuffled), 52)
	assert.Equal(t, len(deck), len(shuffled))
	matchCount := 52
	for i := range deck {
		if deck[i] == shuffled[i] {
			matchCount--
		}
	}
	if matchCount == 0 {
		t.Errorf("Deck not shuffled")
	}
}

func TestResetDeck(t *testing.T) {
	deck := NewDeck()
	assert.Equal(t, len(deck), 52)
	deck = deck[:0]
	assert.Equal(t, len(deck), 0)
	deck.Reset()
	assert.Equal(t, len(deck), 52)
}

func TestDrawCard(t *testing.T) {
	deck := NewDeck()
	deck.Shuffle()

	assert.Equal(t, len(deck), 52)

	c, _ := deck.Draw()
	assert.Equal(t, len(deck), 51)
	assert.Equal(t, reflect.TypeOf(c).String(), "playingcards.Card")
}

func TestDrawCardError(t *testing.T) {
	deck := NewDeck()
	deck.Shuffle()
	assert.Equal(t, len(deck), 52)
	for i := 0; i < 52; i++ {
		_, _ = deck.Draw()
	}
	_, err := deck.Draw()
	assert.Equal(t, err.Error(), ErrNoCardsLeft.Error())
}

func TestDrawN(t *testing.T) {
	tests := []struct {
		name        string
		cardsToDraw int
		expectError bool
	}{
		{
			name:        "Draw 0 cards",
			cardsToDraw: 0,
			expectError: false,
		},
		{
			name:        "Draw 5 cards",
			cardsToDraw: 5,
			expectError: false,
		},
		{
			name:        "Draw 52 cards",
			cardsToDraw: 52,
			expectError: false,
		},
		{
			name:        "Draw 53 cards",
			cardsToDraw: 53,
			expectError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deck := NewDeck()
			deck.Shuffle()

			assert.Equal(t, len(deck), 52)
			cards, err := deck.DrawN(tt.cardsToDraw)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NotError(t, err)
				assert.Equal(t, len(cards), tt.cardsToDraw)
				assert.Equal(t, len(deck), 52-tt.cardsToDraw)
			}
		})
	}
}

func TestCardsLeft(t *testing.T) {
	tests := []struct {
		name        string
		cardsToDraw int
		expectError bool
	}{
		{
			name:        "Draw 0 cards",
			cardsToDraw: 0,
			expectError: false,
		},
		{
			name:        "Draw 5 cards",
			cardsToDraw: 5,
			expectError: false,
		},
		{
			name:        "Draw 52 cards",
			cardsToDraw: 52,
			expectError: false,
		},
		{
			name:        "Draw 53 cards",
			cardsToDraw: 53,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deck := NewDeck()
			deck.Shuffle()
			assert.Equal(t, len(deck), 52)
			cards, err := deck.DrawN(tt.cardsToDraw)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NotError(t, err)
				assert.Equal(t, len(cards), tt.cardsToDraw)
				assert.Equal(t, deck.CardsLeft(), 52-tt.cardsToDraw)
			}
		})
	}
}
