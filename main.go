// Smoke Or Fire. The card game introduced by Kyle Durham. Coded into digital form by Matthew Brandeburg in January 2019

package main

import (
	"github.com/gophercises/deck"	
	"strings"
	"time"
	"fmt"
	_ "golang.org/x/mobile/app" // for the mobile version
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
		fmt.Println("\nSecond, how many decks are we playing with? (1-20)")
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
	fmt.Println("\nLastly, for Player1, my question is...")
	
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
				fmt.Printf("Player%d had a tie with %s and is safe!\n", p.Number, card)
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
	fmt.Println("***********************************************************")
	fmt.Println("***********************************************************")
	fmt.Println("***********************************************************")
	time.Sleep(1 * time.Second)	
	fmt.Println("\nNow begins the real game... the good, the bad, and the ugly\n")
	time.Sleep(1 * time.Second)
	

	fmt.Println("To recap, the players have the following hands:")
	for _, p := range players {
		fmt.Printf("\nPlayer%d: %s, %s, %s, %s\n", p.Number, p.Hand[0], p.Hand[1], p.Hand[2], p.Hand[3])
	}
	time.Sleep(5 * time.Second)


	// need a for loop to play through the remainder of the cards
	// fmt.Println(len(cards)) //For comparison, remember, there should be 44 (52-8)")
	newLen := len(cards) % 3
	if newLen == 1 { // and if we play the games in these, then we dont have to worry about calling them back later
		fmt.Println("\nThere will be one bonus ugly card!\n")
		ugly1 :=  cards[len(cards)-1]
		cards = cards[:len(cards)-1]

		var goodStack []deck.Card
		var badStack []deck.Card
		var uglyStack []deck.Card

		fmt.Printf("Ready to deal the first card? (Hit enter)\n") //I think this can be a poor man's break as we do all the next rounds
		fmt.Scanf("%s\n", &input)
		deckLength := len(cards)
		for i := 1; i <= deckLength; i++ {
			if i % 3 == 1 {
				// Good
				card, cards = draw(cards)
				goodStack = append(goodStack, card)

				for _, p := range players {
					fmt.Printf("\nPlayer%d: %s, %s, %s, %s\n", p.Number, p.Hand[0], p.Hand[1], p.Hand[2], p.Hand[3])
				}
				fmt.Println("\nGOOD!")
				fmt.Printf("...%s!\n", card)

				for _, p := range players {
					for z := 0; z < 4; z++ {
						if int(p.Hand[z].Rank) == int(card.Rank) {
							fmt.Printf("\nPlayer%d matched with the %s (%s) and gets to give a drink of one second!\n", p.Number, card, p.Hand[z])
						} 
					}
				}

				fmt.Printf("\nReady to for the next round? (Hit enter)\n")
				fmt.Scanf("%s\n", &input) // again, doesn't do anything, but gives us time before moving on


				} else if i % 3 == 2 {
					// Bad
					card, cards = draw(cards)
					badStack = append(badStack, card)

					for _, p := range players {
						fmt.Printf("\nPlayer%d: %s, %s, %s, %s\n", p.Number, p.Hand[0], p.Hand[1], p.Hand[2], p.Hand[3])
					}
					fmt.Println("\nBAD!")
					fmt.Printf("...%s!\n", card)

					for _, p := range players {
						for z := 0; z < 4; z++ {
							if int(p.Hand[z].Rank) == int(card.Rank) {
								fmt.Printf("\nPlayer%d matched with the %s (%s) and has to drink for one second!\n", p.Number, card, p.Hand[z])
							} 
						}
					}

					fmt.Printf("\nReady to for the next round? (Hit enter)\n")
					fmt.Scanf("%s\n", &input) // again, doesn't do anything, but gives us time before moving on

					} else if i % 3 == 0 {
						// Ugly
						card, cards = draw(cards)
						uglyStack = append(uglyStack, card)

						for _, p := range players {
							fmt.Printf("\nPlayer%d: %s, %s, %s, %s\n", p.Number, p.Hand[0], p.Hand[1], p.Hand[2], p.Hand[3])
						}
						fmt.Println("\nUGLY!")
						fmt.Printf("...%s!\n", card)

						for _, p := range players {
							for z := 0; z < 4; z++ {
								if int(p.Hand[z].Rank) == int(card.Rank) {
									fmt.Printf("\nPlayer%d matched with the %s (%s) and has to drink for %d seconds!\n", p.Number, card, p.Hand[z], int(card.Rank))
								} 
							}
						}

						fmt.Printf("\nReady to for the next round? (Hit enter)\n")
						fmt.Scanf("%s\n", &input) // again, doesn't do anything, but gives us time before moving on

					}
		}

		fmt.Println("\nLastly, now for the bonus ugly card!\n")
		time.Sleep(1 * time.Second)
		// fmt.Println(ugly1)
		fmt.Printf("...%s!\n", ugly1)
		// Player hand scan function (will need for matching int(card.Rank) for drink hits)
		for _, p := range players {
			for i := 0; i < 4; i++ {
				if int(p.Hand[i].Rank) == int(ugly1.Rank) {
					fmt.Printf("Player%d matched with the %s and has to drink for %d seconds!\n", p.Number, ugly1, int(ugly1.Rank))
					} 
				} 
			}
			// End of player scan function

	} else if newLen == 2 {
		fmt.Println("\nThere will be two bonus ugly cards!\n")
		
		ugly1 := cards[len(cards)-1]
		cards = cards[:len(cards)-1]

		ugly2 := cards[len(cards)-1]
		cards = cards[:len(cards)-1] // print(len(cards)) // should now be 42 (edit: and it is!)

		/*  code what happens below for main gameplay (testing if you hadn't noticed on 2 players, 1 deck configuration) */

		var goodStack []deck.Card
		var badStack []deck.Card
		var uglyStack []deck.Card

		fmt.Printf("Ready to deal the first card? (Hit enter)\n") //I think this can be a poor man's break as we do all the next rounds
		fmt.Scanf("%s\n", &input)
		deckLength := len(cards)
		for i := 1; i <= deckLength; i++ {
			if i % 3 == 1 {
				// Good
				card, cards = draw(cards)
				goodStack = append(goodStack, card)

				for _, p := range players {
					fmt.Printf("\nPlayer%d: %s, %s, %s, %s\n", p.Number, p.Hand[0], p.Hand[1], p.Hand[2], p.Hand[3])
				}
				fmt.Println("\nGOOD!")
				fmt.Printf("...%s!\n", card)

				for _, p := range players {
					for z := 0; z < 4; z++ {
						if int(p.Hand[z].Rank) == int(card.Rank) {
							fmt.Printf("\nPlayer%d matched with the %s (%s) and gets to give a drink of one second!\n", p.Number, card, p.Hand[z])
						} 
					}
				}

				fmt.Printf("\nReady to for the next round? (Hit enter)\n")
				fmt.Scanf("%s\n", &input) // again, doesn't do anything, but gives us time before moving on


				} else if i % 3 == 2 {
					// Bad
					card, cards = draw(cards)
					badStack = append(badStack, card)

					for _, p := range players {
						fmt.Printf("\nPlayer%d: %s, %s, %s, %s\n", p.Number, p.Hand[0], p.Hand[1], p.Hand[2], p.Hand[3])
					}
					fmt.Println("\nBAD!")
					fmt.Printf("...%s!\n", card)

					for _, p := range players {
						for z := 0; z < 4; z++ {
							if int(p.Hand[z].Rank) == int(card.Rank) {
								fmt.Printf("\nPlayer%d matched with the %s (%s) and has to drink for one second!\n", p.Number, card, p.Hand[z])
							} 
						}
					}

					fmt.Printf("\nReady to for the next round? (Hit enter)\n")
					fmt.Scanf("%s\n", &input) // again, doesn't do anything, but gives us time before moving on

					} else if i % 3 == 0 {
						// Ugly
						card, cards = draw(cards)
						uglyStack = append(uglyStack, card)

						for _, p := range players {
							fmt.Printf("\nPlayer%d: %s, %s, %s, %s\n", p.Number, p.Hand[0], p.Hand[1], p.Hand[2], p.Hand[3])
						}
						fmt.Println("\nUGLY!")
						fmt.Printf("...%s!\n", card)

						for _, p := range players {
							for z := 0; z < 4; z++ {
								if int(p.Hand[z].Rank) == int(card.Rank) {
									fmt.Printf("\nPlayer%d matched with the %s (%s) and has to drink for %d seconds!\n", p.Number, card, p.Hand[z], int(card.Rank))
								} 
							}
						}

						fmt.Printf("\nReady to for the next round? (Hit enter)\n")
						fmt.Scanf("%s\n", &input) // again, doesn't do anything, but gives us time before moving on

					}
		}
		// fmt.Printf("To test, the number of cards we didn't deal out, minus bonus uglies, are: %d", len(cards))
		// ********************* End of coding main gameplay *************************************************

		fmt.Println("\nLastly, now for the last two bonus ugly cards!\n")
		time.Sleep(1 * time.Second)
		// fmt.Println(int(ugly1.Rank))
		// fmt.Println(int(ugly2.Rank))
		fmt.Printf("...%s and %s!\n", ugly1, ugly2)
		for _, p := range players {
			for i := 0; i < 4; i++ {
				if int(p.Hand[i].Rank) == int(ugly1.Rank) {
					fmt.Printf("Player%d matched with the %s and has to drink for %d seconds!\n", p.Number, ugly1, int(ugly1.Rank))
					} else if int(p.Hand[i].Rank) == int(ugly2.Rank) {
						fmt.Printf("Player%d matched with the %s and has to drink for %d seconds!\n", p.Number, ugly2, int(ugly1.Rank))
						} 
					}
			}
	} else {
		fmt.Println("\nThere won't be any bonus ugly cards!\n")
		
		var goodStack []deck.Card
		var badStack []deck.Card
		var uglyStack []deck.Card

		fmt.Printf("Ready to deal the first card? (Hit enter)\n") //I think this can be a poor man's break as we do all the next rounds
		fmt.Scanf("%s\n", &input)
		deckLength := len(cards)
		for i := 1; i <= deckLength; i++ {
			if i % 3 == 1 {
				// Good
				card, cards = draw(cards)
				goodStack = append(goodStack, card)

				for _, p := range players {
					fmt.Printf("\nPlayer%d: %s, %s, %s, %s\n", p.Number, p.Hand[0], p.Hand[1], p.Hand[2], p.Hand[3])
				}
				fmt.Println("\nGOOD!")
				fmt.Printf("...%s!\n", card)

				for _, p := range players {
					for z := 0; z < 4; z++ {
						if int(p.Hand[z].Rank) == int(card.Rank) {
							fmt.Printf("\nPlayer%d matched with the %s (%s) and gets to give a drink of one second!\n", p.Number, card, p.Hand[z])
						} 
					}
				}

				fmt.Printf("\nReady to for the next round? (Hit enter)\n")
				fmt.Scanf("%s\n", &input) // again, doesn't do anything, but gives us time before moving on


				} else if i % 3 == 2 {
					// Bad
					card, cards = draw(cards)
					badStack = append(badStack, card)

					for _, p := range players {
						fmt.Printf("\nPlayer%d: %s, %s, %s, %s\n", p.Number, p.Hand[0], p.Hand[1], p.Hand[2], p.Hand[3])
					}
					fmt.Println("\nBAD!")
					fmt.Printf("...%s!\n", card)

					for _, p := range players {
						for z := 0; z < 4; z++ {
							if int(p.Hand[z].Rank) == int(card.Rank) {
								fmt.Printf("\nPlayer%d matched with the %s (%s) and has to drink for one second!\n", p.Number, card, p.Hand[z])
							} 
						}
					}

					fmt.Printf("\nReady to for the next round? (Hit enter)\n")
					fmt.Scanf("%s\n", &input) // again, doesn't do anything, but gives us time before moving on

					} else if i % 3 == 0 {
						// Ugly
						card, cards = draw(cards)
						uglyStack = append(uglyStack, card)

						for _, p := range players {
							fmt.Printf("\nPlayer%d: %s, %s, %s, %s\n", p.Number, p.Hand[0], p.Hand[1], p.Hand[2], p.Hand[3])
						}
						fmt.Println("\nUGLY!")
						fmt.Printf("...%s!\n", card)

						for _, p := range players {
							for z := 0; z < 4; z++ {
								if int(p.Hand[z].Rank) == int(card.Rank) {
									fmt.Printf("\nPlayer%d matched with the %s (%s) and has to drink for %d seconds!\n", p.Number, card, p.Hand[z], int(card.Rank))
								} 
							}
						}

						fmt.Printf("\nReady to for the next round? (Hit enter)\n")
						fmt.Scanf("%s\n", &input) // again, doesn't do anything, but gives us time before moving on

					}
		}

	}
	fmt.Printf("\nThanks for playing! (Hit enter to quit)\n")
	fmt.Scanf("%s\n", &input) // again, doesn't do anything, but gives us time before moving on
}

// env GOOS=windows GOARCH=386 go build -v github.com/gophercises/smokeOrFire
//$ gomobile build -target=ios -bundleid=smokeOrFire github.com/gophercises/smokeOrFire
