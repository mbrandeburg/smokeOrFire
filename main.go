// Smoke Or Fire. The card game introduced by Kyle Durham. Coded into digital form by Matthew Brandeburg in January 2019

package main

// Other Notes:
//$ gomobile build -target=ios -bundleid=smokeOrFire github.com/gophercises/smokeOrFire
//$ git push origin functionalied:master
// fmt.Printf("Player1 drew the %s, so drink for %v seconds!\n", player1[0], int(card.Rank))

import (
	"github.com/gophercises/deck"	
	"strings"
	"time"
	"fmt"
)

type Player struct {
	Number int
	Hand []deck.Card
}

func draw(cards []deck.Card) (deck.Card, []deck.Card){
	return cards[0], cards[1:]	
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}


func main() {
	fmt.Println("\n\n\nWELCOME TO\n")
	time.Sleep(1 * time.Second)
	fmt.Println("SMOKE\n OR\n FIRE\n") //FIRE fire fire fire (whisper) fire
	time.Sleep(1 * time.Second)
	fmt.Println("I have three questions for you..")
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

	var deckCount int
	for {
		fmt.Println("Second, how many decks are we playing with? (1-20)")
		fmt.Scanf("%d\n", &deckCount)
		if deckCount >= 1 && deckCount <= 20 {
			break
		}
		fmt.Println("Please enter a number in the valid range.")
	}
	cards := deck.New(deck.Deck(deckCount), deck.Shuffle)
	var card deck.Card


	// for currPlayer := 0; ; currPlayer = (currPlayer + 1) % len(players){
	time.Sleep(1 * time.Second)
	fmt.Println("Lastly, for Player1, my question is...")
	var input string

	for _, p := range players {
		card, cards = draw(cards)
		p.Hand = append(p.Hand,card)
		time.Sleep(1 * time.Second)
		fmt.Printf("\nPlayer %d: (S)moke or (F)ire?\n", p.Number)
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
	time.Sleep(1 * time.Second) // pause afer the player's round

	for round := 0; round < 3; round ++ {
		if round == 0 {
			for _, p := range players {
			card, cards = draw(cards)
			p.Hand = append(p.Hand,card)
			time.Sleep(1 * time.Second)
			
			prev := p.Hand[round]
			fmt.Printf("\nPlayer %d, last hand you drew a %s, so do you think: (H)igher or (L)ower?\n", p.Number, prev)
			fmt.Scanf("%s\n", &input)

			
			if card.Rank == prev.Rank {
				fmt.Printf("Player%d had a tie with %s this round and %s last round, and is safe!\n", p.Number, card, prev)
				continue
			}

			if strings.ToLower(input) == "h" {
				if card.Rank > prev.Rank {
				fmt.Printf("Player%d had a higher card with %s than his or her previous card of %s and is safe!\n", p.Number, card, prev)
					} else {
					fmt.Printf("Player%d had a lower card with %s than his or her previous card of %s and has to drink for one second.\n", p.Number, card, prev)
				} } else { // they chose lower
					if card.Rank < prev.Rank {
						fmt.Printf("Player%d had a lower card with %s than his or her previous card of %s and is safe!\n", p.Number, card, prev)
					} else {
						fmt.Printf("Player%d had a higher card with %s than his or her previous card of %s and has to drink for one second.\n", p.Number, card, prev)
					}
				}
				time.Sleep(1 * time.Second) // pause afer the player's round
			} 
		} else if round == 1 {
			// fmt.Println("This will be the inside/outside round.")
			// next round in the loop! (for _, p := range players ...)
			for _, p := range players {
			card, cards = draw(cards)
			p.Hand = append(p.Hand,card)
			time.Sleep(1 * time.Second)
			
			prev := p.Hand[round]
			fmt.Printf("\nPlayer %d, you've drawn %s and %s, so do you think: (I)nside or (O)utside?\n", p.Number, p.Hand[0], prev)
			fmt.Scanf("%s\n", &input)
			
			// need to determine max and min of p.Hand for inside or outside, since could have come in any order during the previous rounds
			min1 := min(int(p.Hand[0].Rank),int(p.Hand[1].Rank))
			max1 := max(int(p.Hand[0].Rank),int(p.Hand[1].Rank))
			// fmt.Println(min1,max1) // it works!
			
			if int(card.Rank) == min1 || int(card.Rank) == max1 {
				fmt.Printf("Player%d had a tie and is safe!\n", p.Number)
				continue
			}

			if strings.ToLower(input) == "i" { // I need to make a minimum and a maximum function
				if int(card.Rank) > min1 || int(card.Rank) < max1 {
				fmt.Printf("Player%d's card of %s is inside and is safe!\n", p.Number, card)
					} else {
					fmt.Printf("Player%d's card of %s is outside and he or she has to drink for one second.\n", p.Number, card)
				} } else { // they chose outside
					if int(card.Rank) > max1 || int(card.Rank) < min1 {
						fmt.Printf("Player%d's card of %s is outside and is safe!\n", p.Number, card)
					} else {
						fmt.Printf("Player%d's card of %s is inside and he or she has to drink for one second.\n", p.Number, card)
					}
				}
			}
			time.Sleep(1 * time.Second) // pause afer the player's round

		} else if round == 2 {
			// fmt.Println("This will be the odd/even round.")
			for _, p := range players {
			card, cards = draw(cards)
			p.Hand = append(p.Hand,card)
			time.Sleep(1 * time.Second)
			
			fmt.Printf("\nPlayer %d, do you think: (O)dd or (E)ven?\n", p.Number)
			fmt.Scanf("%s\n", &input)
			
			if strings.ToLower(input) == "e" {
				if int(card.Rank) % 2 == 0 {
				fmt.Printf("Player%d's card of %s is even and is safe!\n", p.Number, card)
					} else {
					fmt.Printf("Player%d's card of %s is odd and he or she has to drink for one second.\n", p.Number, card)
				} } else { // they chose odd
					if int(card.Rank) % 2 == 0 {
						fmt.Printf("Player%d's card of %s is even and he or she has to drink for one second.\n", p.Number, card)
					} else {
						fmt.Printf("Player%d's card of %s is odd and is safe!\n", p.Number, card)
					}
				}
			}
		}
		time.Sleep(1 * time.Second) // pause afer the player's round
	}		
	fmt.Println("Now begins the real game... the good, the bad, and the ugly")
	time.Sleep(1 * time.Second)
	fmt.Println("\n")
	// Now for everything else I guess...
}
