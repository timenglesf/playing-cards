package playingcards

import (
	"testing"

	"github.com/timenglesf/playing-cards/internal/assert"
)

func TestSuitString(t *testing.T) {
	tests := []struct {
		suit Suit
		want string
	}{
		{Diamonds, "Diamonds"},
		{Clubs, "Clubs"},
		{Hearts, "Hearts"},
		{Spades, "Spades"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.suit.String(), tt.want)
	}
}

func TestRankString(t *testing.T) {
	tests := []struct {
		rank Rank
		want string
	}{
		{Ace, "Ace"},
		{Two, "Two"},
		{Three, "Three"},
		{Four, "Four"},
		{Five, "Five"},
		{Six, "Six"},
		{Seven, "Seven"},
		{Eight, "Eight"},
		{Nine, "Nine"},
		{Ten, "Ten"},
		{Jack, "Jack"},
		{Queen, "Queen"},
		{King, "King"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.rank.String(), tt.want)
	}
}

func TestCardString(t *testing.T) {
	tests := []struct {
		card Card
		want string
	}{
		{
			card: Card{Suit: Diamonds, Rank: Ace},
			want: "Ace of Diamonds",
		},
		{
			card: Card{Suit: Clubs, Rank: Two},
			want: "Two of Clubs",
		},
		{
			card: Card{Suit: Hearts, Rank: Three},
			want: "Three of Hearts",
		},
		{
			card: Card{Suit: Spades, Rank: Four},
			want: "Four of Spades",
		},
		{
			card: Card{Suit: Diamonds, Rank: Five},
			want: "Five of Diamonds",
		},
		{
			card: Card{Suit: Clubs, Rank: Six},
			want: "Six of Clubs",
		},
		{
			card: Card{Suit: Hearts, Rank: Seven},
			want: "Seven of Hearts",
		},
		{
			card: Card{Suit: Spades, Rank: Eight},
			want: "Eight of Spades",
		},
		{
			card: Card{Suit: Diamonds, Rank: Nine},
			want: "Nine of Diamonds",
		},
		{
			card: Card{Suit: Clubs, Rank: Ten},
			want: "Ten of Clubs",
		},
		{
			card: Card{Suit: Hearts, Rank: Jack},
			want: "Jack of Hearts",
		},
		{
			card: Card{Suit: Spades, Rank: Queen},
			want: "Queen of Spades",
		},
		{
			card: Card{Suit: Diamonds, Rank: King},
			want: "King of Diamonds",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.card.String(), tt.want)
	}
}
