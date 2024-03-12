//go:generate stringer -type=Suit,Rank
// go generator is crazy - first lay out your header like above, then type go generate after cmd into folder to auto create these stringers from your CONSTANTS

package deck

import (
	// "stringer" // don't need this b/c of go:generate !
	"fmt"
	"sort"
	"math/rand"
	"time"
) 

// let's make stuff exportable
type Suit uint8 //there arne't that many, so go uint8

const (
	Spade Suit = iota // iota allows us to increment these, so spade is 0, diamond 1, etc. - dont need to repeat it
	Diamond 
	Club
	Heart
	Joker 
)

var suits = [...]Suit{Spade, Diamond, Club, Heart} // for making a deck later, easier if we make an array now - and notice that joker is a special case to be dealt with later


type Rank uint8

const (
	_ Rank = iota // the rank will correspond to face value
	_ // b/c we start at 0 in iota, so we _ it out, then aces are high here so _ again
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
	Ace // Aces high in my deck
)

const ( //easier for building a deck with suits
	minRank = Two
	maxRank = Ace
)

type Card struct {
	Suit
	Rank
}



func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String() //using the stringer now! instead of returning name hardcoded as "Joker"
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String()) // return "Ace of Hearts"
}


// now that we can return one cards, we need to create a deck of cards!
func New(opts ...func([]Card) []Card) []Card { //dont freak out - the first two []Card pertain to the options bit about taking in and returning slices
	// ^^ New(opts) also rules b/c it means we can do stuff like New(Jokers(n)) and make not just a new deck, but one with n number of Jokers in it [See card_test.go]
	var cards []Card

	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	// now let's work with some functional options - see below for CUSTOM sort (instead of default) and shuffle
	for _, opt := range opts {
		cards = opt(cards)
	}

	// and voila your deck-o-cards
	return cards
}


// ***** let's make sort, shuffle, etc options ***********
// need to pre-set this stuff to end with New(DefaultSort)
func absRank(c Card) int {
	return int(c.Suit) * int(maxRank) + int(c.Rank)
}

// FYI - this comes from "sort" import - the sort package expects a less(i, j int) bool function, and we are building all these following steps ourselves to both understand it better and more easily customize it later
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

// now a generic one for making the customized one
func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card { //we pass in the less function into here
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

///////// SHUFFLE TIME
func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards)) // we will make a new slice for shuffling
// Goal: i.e. rand.Perm(5) should return a slice like [0, 1, 4, 2, 3]
	r := rand.New(rand.NewSource(time.Now().Unix())) //how to make a wildly unique seed - tie it to the date time!
	for i, j := range r.Perm(len(cards)) {
		ret[i] = cards[j]
	}
	return ret
}

// final option - turn on jokers lol
// goal is make fn like New(Jokers(2))
func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Rank: Rank(i), // not traditional ranks, but jokers wild yo!
				Suit: Joker,
			})
		}
		return cards
	}
}


// jk - two more options to go: filtering out some cards, AND ADDING MULTIPLE DECKS!

// filtering - if you want to filter out twos or three for whatever your game is
func Filter(f func(card Card) bool) func([]Card) []Card { //bool bc true or false on if it should filter card or not
	return func(cards []Card) []Card {
		var ret []Card
		for _, c := range cards {
			if !f(c) { // cool way of saying if it returns NOT TRUE...
				ret = append(ret, c)
			}
		}
		return ret
	}
}

// multiple decks - which we are doing by copying the deck x amount of times at first
func Deck(n int) func(cards []Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card // starting with an empty slice, so when we do New(Deck(3)) it acutally copies enough to make 3 decks total of cards, not 4 (see test for verification)
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}
