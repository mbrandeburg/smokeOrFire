package deck

import (
	"fmt"
	"math/rand"
	"time"
)

// Suit represents a card suit
type Suit int

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
)

func (s Suit) String() string {
	return [...]string{"Spades", "Diamonds", "Clubs", "Hearts"}[s]
}

// Rank represents a card rank
type Rank int

const (
	_ Rank = iota
	Ace
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

func (r Rank) String() string {
	return [...]string{"", "Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}[r]
}

// Card represents a playing card
type Card struct {
	Suit Suit `json:"suit"`
	Rank Rank `json:"rank"`
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank, c.Suit)
}

// New creates a new deck of cards
func New(options ...func([]Card) []Card) []Card {
	var cards []Card
	
	// Create standard 52-card deck
	for suit := Spade; suit <= Heart; suit++ {
		for rank := Ace; rank <= King; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	
	// Apply options
	for _, option := range options {
		cards = option(cards)
	}
	
	return cards
}

// Shuffle is an option function to shuffle the deck
func Shuffle(cards []Card) []Card {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(cards) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
	return cards
}

// Deck is an option function to multiply the deck
func Deck(count int) func([]Card) []Card {
	return func(cards []Card) []Card {
		if count <= 1 {
			return cards
		}
		
		original := make([]Card, len(cards))
		copy(original, cards)
		
		for i := 1; i < count; i++ {
			cards = append(cards, original...)
		}
		
		return cards
	}
}
