package main

//$ gomobile build -target=ios -bundleid=smokeOrFire golang.org/x/mobile/example/basic

import (
	"github.com/gophercises/deck"	
	// "strings"
	"fmt"
)


func draw(cards []deck.Card) (deck.Card, []deck.Card){
	return cards[0], cards[1:]	
}

func main() {
	cards := deck.New(deck.Shuffle)
	var card deck.Card
	for i := 0; i < 52; i++ { // not <= b/c there's a 0th card
		card, cards = draw(cards)
		fmt.Println(card)
		// fmt.Println(i)
	}
	// fmt.Println(len(cards))
}


///////
// func main() {
// 	cards := deck.New(deck.Shuffle)
// 	var card deck.Card
// 	var player1, player2 Hand
// 	for i := 0; i < len(cards); i++ { // do it twice so they each get two cards
// 		for _, hand := range []*Hand{&player1, &player2} { // use pointers to iterate over them quickly
// 		card, cards = draw(cards)
// 		*hand = append(*hand, card) //pointer here means updating the VALUE that it points to
// 		}
// 		fmt.Println("Player One:", player1)
// 		fmt.Printf("Drink for %d seconds!\n", player1.Score())
// 		fmt.Println("Player Two:", player2)
// 		fmt.Printf("Drink for %v seconds!\n", player2.Score())
// 	}
// }
