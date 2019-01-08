package main

//$ gomobile build -target=ios -bundleid=smokeOrFire golang.org/x/mobile/example/basic

import (
	"github.com/gophercises/deck"	
	"strings"
	"fmt"
)

// func main() {
// 	cards := deck.New(deck.Shuffle)
// 	var card deck.Card
// 	for i := 0; i < 10; i++ { //let's iterate over (aka deal out) 10 cards to start with
// 		card, cards = cards[0], cards[1:] //pulling out the first card, leaving a slice leftover with all the rest of cards till the i++ iterator runs out
// 		fmt.Println(card)
// 	}
// 	fmt.Println("End of 10 card deal.")
// }


// VERSION 2: YOU'RE GIVEN ALL THE CARDS

func main() {
	cards := deck.New(deck.Shuffle)
	var card deck.Card
	// var hand []deck.Card // your hand is a slice of the deck
	for i := 0; i <= len(cards) *2 + 50; i++ { // now why is that???
		card, cards = cards[0], cards[1:]
		fmt.Println(card) // how come this person only gets half? You print out 26 of the 52 cards
		// hand = append(hand, card)
		}
	// fmt.Println("Player 1:", hand)
	fmt.Println("And the rest are:")
	fmt.Println(cards)
}


// make a player hand
type Hand []deck.Card 
func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h { // i is that first, so its index (if we did , smthg then smthg would be value)
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

