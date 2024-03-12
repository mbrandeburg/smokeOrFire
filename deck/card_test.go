// Test with expected output!!!
package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade}) 
	fmt.Println(Card{Suit: Joker})
	
	// Output: 
	// Ace of Hearts
	// Two of Spades
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Error("Wrong number of cards in a new deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Two, Suit: Spade}
	if cards[0] != exp { 
		t.Error("Expected Two of Spades as first card, but received", cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(2))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 2 {
		t.Error("Expected 2 Jokers, received:", count)
	}
}


func TestFilter(t *testing.T) {
	// build the filter function here that will throw out twos and threes for whatever reason
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected all twos and threes to be filtered out, but woops.")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	if len(cards) != 52*3 {
		t.Errorf("Expected %d cards, received %d cards.", 52*3, len(cards))
	}
}