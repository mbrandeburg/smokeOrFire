package main

//$ gomobile build -target=ios -bundleid=smokeOrFire golang.org/x/mobile/example/basic

import (
	"github.com/gophercises/deck"	
	// "strings"
	"fmt"
)


// make a series of switches - so you base it on number of players? if 2, then trigger draw function for two players, etc.
// make it custotmizable up to 8 players



// what if we try this without the iterator or hands
func draw(cards []deck.Card) (deck.Card, []deck.Card){
	return cards[0], cards[1:]	
}


func (c Card) Score() int {
	score := 0
	for _, c := range h {
		score += int(c.Rank) // see below - need to keep facecards from going over 10
	}
	return score
}

func main() {
	cards := deck.New(deck.Shuffle)
	var card deck.Card
	var player1, player2 cards

	// need to move draw funciton to the player bucket


	fmt.Println("Player One:", player1)
	fmt.Printf("Drink for %d seconds!\n", player1.Score())
	fmt.Println("Player Two:", player2)
	fmt.Printf("Drink for %v seconds!\n", player2.Score())
}


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

// func main() {
// 	cards := deck.New(deck.Shuffle)
// 	var card deck.Card
// 	// var hand []deck.Card // your hand is a slice of the deck
// 	for i := 0; i <= len(cards) *2 + 50; i++ { // now why is that???
// 		card, cards = cards[0], cards[1:]
// 		fmt.Println(card) // how come this person only gets half? You print out 26 of the 52 cards
// 		// hand = append(hand, card)
// 		}
// 	// fmt.Println("Player 1:", hand)
// 	fmt.Println("And the rest are:")
// 	fmt.Println(cards)
// }


// // make a player hand
// type Hand []deck.Card 
// func (h Hand) String() string {
// 	strs := make([]string, len(h))
// 	for i := range h { // i is that first, so its index (if we did , smthg then smthg would be value)
// 		strs[i] = h[i].String()
// 	}
// 	return strings.Join(strs, ", ")
// }

// Version 3: let's work on scoring
// func (h Hand) Score() int {
// 	score := 0
// 	for _, c := range h {
// 		score += int(c.Rank) // see below - need to keep facecards from going over 10
// 	}
// 	return score
// }

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
