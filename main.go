package main

import (
	"github.com/gophercises/deck"	
	"fmt"
)

func main() {
	cards := deck.New(deck.Shuffle)
	var card deck.Card
	for i := 0; i < 10; i++ { //let's iterate over (aka deal out) 10 cards to start with
		card, cards = cards[0], cards[1:] //pulling out the first card, leaving a slice leftover with all the rest of cards till the i++ iterator runs out
		fmt.Println(card)
	}
	fmt.Println("End of 10 card deal.")
}