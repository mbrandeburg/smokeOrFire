package main

//$ gomobile build -target=ios -bundleid=smokeOrFire golang.org/x/mobile/example/basic

import (
	"github.com/gophercises/deck"	
	"strings"
	"time"
	"fmt"
)


// fmt.Printf("Player1 drew the %s, so drink for %v seconds!\n", player1[0], int(card.Rank))

func draw(cards []deck.Card) (deck.Card, []deck.Card){
	return cards[0], cards[1:]	
}

// func score(cards []deck.Card) int {
// 	score := cards.Rank
// 	return score
// }

type Player struct {
	Number int
	Hand []deck.Card
}

func main() {
	cards := deck.New(deck.Shuffle)
	var card deck.Card
	fmt.Println("\n\n\nWELCOME TO\n")
	fmt.Println("SMOKE\n OR\n FIRE\n") //FIRE fire fire fire (whisper) fire
	fmt.Println("I have two questions for you..")

	var playerCount int
	for {
		fmt.Println("How many players are there? (2-8)")
		fmt.Scanf("%d\n", &playerCount)
		if playerCount >= 2 && playerCount <= 8 {
			break
		}
	}

	var players []*Player
	for i := 0; i < playerCount; i++ {
		players = append(players, &Player {
			Number: i, 
			})
		}

	// for currPlayer := 0; ; currPlayer = (currPlayer + 1) % len(players){
	
	for _, p := range players {
		card, cards = draw(cards)
		p.Hand = append(p.Hand,card)
		var input string
		time.Sleep(1 * time.Second)
		fmt.Printf("Player %d: (S)moke or (F)ire?\n", p.Number)
		fmt.Scanf("%s\n", &input)
		if strings.ToLower(input) == "s" {
			switch card.Suit {
			case deck.Spade, deck.Club:
				// safe
			case deck.Diamond, deck.Heart:
				//oops
			}

		} else {

		}

	}
	for round := 0; round < 3; round ++ {
		for _, p := range players {
		card, cards = draw(cards)
		p.Hand = append(p.Hand,card)
		
		time.Sleep(1 * time.Second)
		fmt.Printf("Player %d: (H)igher or (L)ower?\n", p.Number)
		fmt.Scanf("%s\n", &input)

		prev := p.Hand[round]
		if card.Rank == prev.Rank {
			//safe
			continue
		}

		if strings.ToLower(input) == "h" {
			card.Rank > prev.Rank {
			// safe
			continue
			} else {

		}

	}		
	// 

}

	



		switch playerCount {
		case 2:
			var player1 []deck.Card
			var player2 []deck.Card
			var input string

			fmt.Println("...and...for player1...")
			time.Sleep(1 * time.Second)
			fmt.Println("(S)moke or (F)ire?")
			fmt.Scanf("%s\n", &input)
				switch input {
				case "s":
					card, cards = draw(cards)
					player1 = append(player1, card)
					if card.Suit.String() == "Spade" || card.Suit.String() == "Club" {
						fmt.Printf("Player1 drew the %s and is safe this round.\n", player1[0])
					} else {
						fmt.Printf("Player1 drew the %s and has to drink for one second.\n", player1[0])
					}
				case "f":
					card, cards = draw(cards)
					player1 = append(player1, card)
					if card.Suit.String() == "Diamond" || card.Suit.String() == "Heart" {
						fmt.Printf("Player1 drew the %s and is safe this round.\n", player1[0])
					} else {
						fmt.Printf("Player1 drew the %s and has to drink for one second.\n", player1[0])
					}
				}
			time.Sleep(2 * time.Second) // take a pause between turns!
			fmt.Println("Player2, which do you chose?")
			fmt.Println("(S)moke or (F)ire?")
			fmt.Scanf("%s\n", &input)
				switch input {
				case "s":
					card, cards = draw(cards)
					player2 = append(player2, card)
					if card.Suit.String() == "Spade" || card.Suit.String() == "Club" {
						fmt.Printf("Player1 drew the %s and is safe this round.\n", player2[0])
					} else {
						fmt.Printf("Player1 drew the %s and has to drink for one second.\n", player2[0])
					}
				case "f":
					card, cards = draw(cards)
					player2 = append(player2, card)
					if card.Suit.String() == "Diamond" || card.Suit.String() == "Heart" {
						fmt.Printf("Player1 drew the %s and is safe this round.\n", player2[0])
					} else {
						fmt.Printf("Player1 drew the %s and has to drink for one second.\n", player2[0])
					}
				}
			time.Sleep(1 * time.Second) // pause before next question series
		// case 3:
		// 	var player1 []deck.Card
		// 	var player2 []deck.Card
		// 	var player3 []deck.Card
		// case 4:
		// 	var player1 []deck.Card
		// 	var player2 []deck.Card
		// 	var player3 []deck.Card
		// 	var player4 []deck.Card
		// case 5:
		// 	var player1 []deck.Card
		// 	var player2 []deck.Card
		// 	var player3 []deck.Card
		// 	var player4 []deck.Card
		// 	var player5 []deck.Card
		// case 6:
		// 	var player1 []deck.Card
		// 	var player2 []deck.Card
		// 	var player3 []deck.Card
		// 	var player4 []deck.Card
		// 	var player5 []deck.Card
		// 	var player6 []deck.Card
		// case 7:
		// 	var player1 []deck.Card
		// 	var player2 []deck.Card
		// 	var player3 []deck.Card
		// 	var player4 []deck.Card
		// 	var player5 []deck.Card
		// 	var player6 []deck.Card
		// 	var player7 []deck.Card
		// case 8:
		// 	var player1 []deck.Card
		// 	var player2 []deck.Card
		// 	var player3 []deck.Card
		// 	var player4 []deck.Card
		// 	var player5 []deck.Card
		// 	var player6 []deck.Card
		// 	var player7 []deck.Card
		// 	var player8 []deck.Card
		default:
			fmt.Println("Invalid number of players. Sorry.")
			// for Chris - is there a way to hit the default case on a switch but try again?
		}
////////// CAN I JUST MAKE IT DEPEND ON CASE HERE? OR IS IT EASIER TO CUT AND PASTE INTO EACH CASE?
	// fmt.Println("...and...for player1...")

	// fmt.Println("(S)moke or (F)ire?")
	// var input string
	// fmt.Scanf("%s\n", &input)
	// 	switch input {
	// 	case "s":
	// 		card, cards = draw(cards)
	// 		// score := int(card.Rank)
	// 		// fmt.Println(score)
	// 		player1 = append(player1, card)
	// 		// fmt.Println(player1)
	// 		fmt.Printf("Player1 drew the %s, so drink for %v seconds!\n", player1[0], int(card.Rank)) //.Score()
	// 	case "f":
	// 		card, cards = draw(cards)
	// 		player1 = append(player1, card)
	// 		fmt.Printf("Player1 drew the %s, so drink for %v seconds!\n", player1[0], int(card.Rank)) //.Score()
	// 	}
}




/////// THE LEFTOVERS

	// for i := 0; i < 2; i++ { // not <=52 b/c there's a 0th card
	// 	card, cards = draw(cards)
	// 	fmt.Println(card)
	// 	// fmt.Println(i)
	// }
	// // fmt.Println(len(cards))


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
