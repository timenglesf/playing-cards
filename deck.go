package playingcards

import (
	"errors"
	"math/rand"
)

// Deck represents a deck of cards.
type Deck []Card

// DeckError represents an error related to a Deck.
type DeckError error

var (
	ErrNoCardsLeft    = DeckError(errors.New("no cards left"))
	ErrNotEnoughCards = DeckError(errors.New("not enough cards"))
)

// NewDeck creates a new deck of 52 cards.
func NewDeck() Deck {
	deck := make(Deck, 0, 52)
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, Card{Suit: suit, Rank: rank})
		}
	}
	return deck
}

// Shuffle shuffles the cards in the deck.
func (d *Deck) Shuffle() {
	rand.Shuffle(len(*d), func(i, j int) {
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	})
}

// Reshuffle shuffles the cards in the deck.
func (d *Deck) Reshuffle() {
	d.Shuffle()
}

// Reset resets the deck to its original state.
func (d *Deck) Reset() {
	*d = NewDeck()
}

// Draw draws a card from the deck. It returns an error if there are no cards left.
func (d *Deck) Draw() (Card, error) {
	if len(*d) == 0 {
		return Card{}, ErrNoCardsLeft
	}
	card := (*d)[0]
	*d = (*d)[1:]
	return card, nil
}

// DrawN draws n cards from the deck. It returns an error if there are not enough cards left.
func (d *Deck) DrawN(n int) ([]Card, error) {
	if len(*d) == 0 {
		return nil, ErrNoCardsLeft
	}
	if len(*d) < n {
		return nil, ErrNotEnoughCards
	}
	cards := make([]Card, n)
	copy(cards, (*d)[:n])
	*d = (*d)[n:]
	return cards, nil
}

// Peek returns the top card in the deck. It returns an error if there are no cards left.
func (d *Deck) Peek() (Card, error) {
	if len(*d) == 0 {
		return Card{}, ErrNoCardsLeft
	}
	return (*d)[0], nil
}

// PeekN returns the top n cards in the deck. It returns an error if there are not enough cards left.
func (d *Deck) PeekN(n int) ([]Card, error) {
	if len(*d) == 0 {
		return nil, ErrNoCardsLeft
	}
	if len(*d) < n {
		return nil, ErrNotEnoughCards
	}
	return (*d)[:n], nil
}

// CardsLeft returns the number of cards left in the deck.
func (d *Deck) CardsLeft() int {
	return len(*d)
}
