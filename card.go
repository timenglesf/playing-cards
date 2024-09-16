package playingcards

import "fmt"

type (
	Suit int
	Rank int
)

type Card struct {
	Suit
	Rank
}

const (
	Diamonds Suit = iota
	Clubs
	Hearts
	Spades
)

const (
	Ace Rank = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

var (
	suits    = []Suit{Diamonds, Clubs, Hearts, Spades}
	suitsStr = [...]string{"Diamonds", "Clubs", "Hearts", "Spades"}
)

var (
	ranks    = []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
	ranksStr = [...]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
)

func (s Suit) String() string {
	return suitsStr[s]
}

func (r Rank) String() string {
	return ranksStr[r]
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank, c.Suit)
}
