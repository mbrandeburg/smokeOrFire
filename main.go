package main

//$ gomobile build -target=ios -bundleid=smokeOrFire golang.org/x/mobile/example/basic

import (
	"github.com/gophercises/deck"	
	"strings"
	"time"
	"fmt"
)


// fmt.Printf("Player1 drew the %s, so drink for %v seconds!\n", player1[0], int(card.Rank))

type Player struct {
	Number int
	Hand []deck.Card
}

func draw(cards []deck.Card) (deck.Card, []deck.Card){
	return cards[0], cards[1:]	
}


func main() {
	cards := deck.New(deck.Shuffle)
	var card deck.Card
	fmt.Println("\n\n\nWELCOME TO\n")
	time.Sleep(1 * time.Second)
	fmt.Println("SMOKE\n OR\n FIRE\n") //FIRE fire fire fire (whisper) fire
	time.Sleep(1 * time.Second)
	fmt.Println("I have two questions for you..")
	time.Sleep(1 * time.Second)

	var playerCount int
	for {
		fmt.Println("First, how many players are there? (2-20)")
		fmt.Scanf("%d\n", &playerCount)
		if playerCount >= 2 && playerCount <= 20 {
			break
		}
		fmt.Println("Please enter a number in the valid range.")
	}

	var players []*Player
	for i := 1; i <= playerCount; i++ { 
		players = append(players, &Player {
			Number: i, 
			})
		}

	// for currPlayer := 0; ; currPlayer = (currPlayer + 1) % len(players){
	time.Sleep(1 * time.Second)
	fmt.Println("Second, for Player1, my question is...")
	for _, p := range players {
		card, cards = draw(cards)
		p.Hand = append(p.Hand,card)
		var input string
		time.Sleep(1 * time.Second)
		fmt.Printf("Player %d: (S)moke or (F)ire?\n", p.Number)
		fmt.Scanf("%s\n", &input)
		if strings.ToLower(input) == "s" {
			switch card.Suit { //already append the hand above
			case deck.Spade, deck.Club: 
				fmt.Printf("Player %d drew the %s and is safe this round.\n", p.Number, p.Hand[0]) //hard coding 0th b/c its the first round
			case deck.Diamond, deck.Heart:
				fmt.Printf("Player %d drew the %s and has to drink for one second.\n", p.Number, p.Hand[0])
				}
			} else { // adding fault tolerance by assuming anything other than smoke will be fire - this "dumb" mechanic repeats throughout
			switch card.Suit { 
			case deck.Diamond, deck.Heart: 
				fmt.Printf("Player %d drew the %s and is safe this round.\n", p.Number, p.Hand[0]) 
			case deck.Spade, deck.Club:
				fmt.Printf("Player %d drew the %s and has to drink for one second.\n", p.Number, p.Hand[0])
			}
		}
	}

	// for round := 0; round < 3; round ++ {
	// 	for _, p := range players {
	// 		card, cards = draw(cards)
	// 		p.Hand = append(p.Hand,card)
			
	// 		time.Sleep(1 * time.Second)
	// 		fmt.Printf("Player %d: (H)igher or (L)ower?\n", p.Number)
	// 		fmt.Scanf("%s\n", &input)

	// 		prev := p.Hand[round]
	// 		if card.Rank == prev.Rank {
	// 			fmt.Printf("Player%d had a tie with %s and is safe!\n", p.Number, p.Hand[round])
	// 			continue
	// 		}

	// 		if strings.ToLower(input) == "h" {
	// 			if card.Rank > prev.Rank {
	// 			fmt.Printf("Player%d had a higher card with %s than his or her previous card of %s and is safe!\n", p.Number, p.Hand[round], prev)
	// 				} else {
	// 				fmt.Printf("Player%d had a lower card with %s than his or her previous card of %s and has to drink for one second.\n", p.Number, p.Hand[round], prev)
	// 			} else { // they chose lower
	// 				if card.Rank < prev.Rank {
	// 					fmt.Printf("Player%d had a lower card with %s than his or her previous card of %s and is safe!\n", p.Number, p.Hand[round], prev)
	// 				} else {
	// 					fmt.Printf("Player%d had a higher card with %s than his or her previous card of %s and has to drink for one second.\n", p.Number, p.Hand[round], prev)
	// 				}
	// 			}
	// 		}
	// 	}
	// 	// next round in the loop! (for _, p := range players ...)
	// }		


	// Now for everything else I guess...

}
