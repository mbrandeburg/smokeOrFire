package game

import (
	"errors"
	"fmt"
	"smokeorfire/pkg/deck"
)

// GamePhase represents the current phase of the game
type GamePhase int

const (
	Setup GamePhase = iota
	SmokeOrFire
	HigherOrLower
	InsideOrOutside
	OddOrEven
	MainGame
	Finished
)

func (p GamePhase) String() string {
	return [...]string{"Setup", "SmokeOrFire", "HigherOrLower", "InsideOrOutside", "OddOrEven", "MainGame", "Finished"}[p]
}

// GuessType represents the type of guess a player can make
type GuessType string

const (
	Smoke   GuessType = "smoke"
	Fire    GuessType = "fire"
	Higher  GuessType = "higher"
	Lower   GuessType = "lower"
	Inside  GuessType = "inside"
	Outside GuessType = "outside"
	Odd     GuessType = "odd"
	Even    GuessType = "even"
)

// Player represents a game player
type Player struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Hand     []deck.Card `json:"hand"`
	IsActive bool        `json:"isActive"`
}

// GameResult represents the result of a player's guess
type GameResult struct {
	PlayerID     int       `json:"playerId"`
	Card         deck.Card `json:"card"`
	Guess        GuessType `json:"guess"`
	IsCorrect    bool      `json:"isCorrect"`
	DrinkSeconds int       `json:"drinkSeconds"`
	Message      string    `json:"message"`
}

// Game represents the game state
type Game struct {
	ID            string       `json:"id"`
	Players       []*Player    `json:"players"`
	Cards         []deck.Card  `json:"-"` // Don't expose the deck
	Phase         GamePhase    `json:"phase"`
	CurrentPlayer int          `json:"currentPlayer"`
	Round         int          `json:"round"`
	GoodStack     []deck.Card  `json:"goodStack"`
	BadStack      []deck.Card  `json:"badStack"`
	UglyStack     []deck.Card  `json:"uglyStack"`
	BonusUgly     []deck.Card  `json:"bonusUgly"`
	Results       []GameResult `json:"results"`
}

// NewGame creates a new game instance
func NewGame(gameID string, playerCount, deckCount int, playerNames []string) (*Game, error) {
	if playerCount < 2 || playerCount > 20 {
		return nil, errors.New("player count must be between 2 and 20")
	}
	if deckCount < 1 || deckCount > 20 {
		return nil, errors.New("deck count must be between 1 and 20")
	}

	// Create players
	players := make([]*Player, playerCount)
	for i := 0; i < playerCount; i++ {
		name := fmt.Sprintf("Player %d", i+1)
		if i < len(playerNames) && playerNames[i] != "" {
			name = playerNames[i]
		}
		
		players[i] = &Player{
			ID:       i + 1,
			Name:     name,
			Hand:     make([]deck.Card, 0, 4),
			IsActive: true,
		}
	}

	// Create and shuffle deck
	cards := deck.New(deck.Deck(deckCount), deck.Shuffle)

	return &Game{
		ID:            gameID,
		Players:       players,
		Cards:         cards,
		Phase:         SmokeOrFire,
		CurrentPlayer: 0,
		Round:         0,
		Results:       make([]GameResult, 0),
	}, nil
}

// Draw draws a card from the deck
func (g *Game) Draw() (deck.Card, error) {
	if len(g.Cards) == 0 {
		return deck.Card{}, errors.New("no more cards in deck")
	}
	
	card := g.Cards[0]
	g.Cards = g.Cards[1:]
	return card, nil
}

// ProcessGuess processes a player's guess and returns the result
func (g *Game) ProcessGuess(playerID int, guess GuessType) (*GameResult, error) {
	if playerID != g.Players[g.CurrentPlayer].ID {
		return nil, errors.New("not this player's turn")
	}

	player := g.Players[g.CurrentPlayer]
	card, err := g.Draw()
	if err != nil {
		return nil, err
	}

	player.Hand = append(player.Hand, card)
	
	result := &GameResult{
		PlayerID: playerID,
		Card:     card,
		Guess:    guess,
	}

	switch g.Phase {
	case SmokeOrFire:
		result = g.processSmokeOrFire(result, card, guess)
	case HigherOrLower:
		result = g.processHigherOrLower(result, card, guess, player)
	case InsideOrOutside:
		result = g.processInsideOrOutside(result, card, guess, player)
	case OddOrEven:
		result = g.processOddOrEven(result, card, guess)
	default:
		return nil, errors.New("invalid game phase for guessing")
	}

	g.Results = append(g.Results, *result)
	g.advancePlayer()
	
	return result, nil
}

func (g *Game) processSmokeOrFire(result *GameResult, card deck.Card, guess GuessType) *GameResult {
	isBlack := card.Suit == deck.Spade || card.Suit == deck.Club
	
	if (guess == Smoke && isBlack) || (guess == Fire && !isBlack) {
		result.IsCorrect = true
		result.DrinkSeconds = 0
		result.Message = fmt.Sprintf("Correct! %s is safe.", result.PlayerID)
	} else {
		result.IsCorrect = false
		result.DrinkSeconds = 1
		result.Message = fmt.Sprintf("Wrong! %s drinks for 1 second.", result.PlayerID)
	}
	
	return result
}

func (g *Game) processHigherOrLower(result *GameResult, card deck.Card, guess GuessType, player *Player) *GameResult {
	prevCard := player.Hand[len(player.Hand)-2] // Previous card
	
	if card.Rank == prevCard.Rank {
		result.IsCorrect = true
		result.DrinkSeconds = 0
		result.Message = "Tie! Safe!"
		return result
	}
	
	if (guess == Higher && card.Rank > prevCard.Rank) || (guess == Lower && card.Rank < prevCard.Rank) {
		result.IsCorrect = true
		result.DrinkSeconds = 0
		result.Message = "Correct! Safe!"
	} else {
		result.IsCorrect = false
		result.DrinkSeconds = 1
		result.Message = "Wrong! Drink for 1 second."
	}
	
	return result
}

func (g *Game) processInsideOrOutside(result *GameResult, card deck.Card, guess GuessType, player *Player) *GameResult {
	// Find min and max of first two cards
	card1, card2 := player.Hand[0], player.Hand[1]
	min := card1.Rank
	max := card2.Rank
	if card2.Rank < card1.Rank {
		min, max = max, min
	}
	
	if card.Rank == min || card.Rank == max {
		result.IsCorrect = true
		result.DrinkSeconds = 0
		result.Message = "Tie! Safe!"
		return result
	}
	
	isInside := card.Rank > min && card.Rank < max
	
	if (guess == Inside && isInside) || (guess == Outside && !isInside) {
		result.IsCorrect = true
		result.DrinkSeconds = 0
		result.Message = "Correct! Safe!"
	} else {
		result.IsCorrect = false
		result.DrinkSeconds = 1
		result.Message = "Wrong! Drink for 1 second."
	}
	
	return result
}

func (g *Game) processOddOrEven(result *GameResult, card deck.Card, guess GuessType) *GameResult {
	isEven := int(card.Rank)%2 == 0
	
	if (guess == Even && isEven) || (guess == Odd && !isEven) {
		result.IsCorrect = true
		result.DrinkSeconds = 0
		result.Message = "Correct! Safe!"
	} else {
		result.IsCorrect = false
		result.DrinkSeconds = 1
		result.Message = "Wrong! Drink for 1 second."
	}
	
	return result
}

func (g *Game) advancePlayer() {
	g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)
	
	// Check if round is complete
	if g.CurrentPlayer == 0 {
		g.Round++
		g.advancePhase()
	}
}

func (g *Game) advancePhase() {
	switch g.Phase {
	case SmokeOrFire:
		g.Phase = HigherOrLower
	case HigherOrLower:
		g.Phase = InsideOrOutside
	case InsideOrOutside:
		g.Phase = OddOrEven
	case OddOrEven:
		g.setupMainGame()
	}
}

func (g *Game) setupMainGame() {
	g.Phase = MainGame
	
	// Calculate bonus ugly cards
	remainingCards := len(g.Cards)
	bonusCount := remainingCards % 3
	
	// Remove bonus ugly cards
	for i := 0; i < bonusCount; i++ {
		card, _ := g.Draw()
		g.BonusUgly = append(g.BonusUgly, card)
	}
}

// ProcessMainGameCard processes the main game phase
func (g *Game) ProcessMainGameCard() (*MainGameResult, error) {
	if g.Phase != MainGame {
		return nil, errors.New("not in main game phase")
	}
	
	if len(g.Cards) == 0 {
		return g.processBonus()
	}
	
	card, err := g.Draw()
	if err != nil {
		return nil, err
	}
	
	cardIndex := len(g.GoodStack) + len(g.BadStack) + len(g.UglyStack)
	var cardType string
	var drinkType string
	
	switch (cardIndex + 1) % 3 {
	case 1: // Good
		g.GoodStack = append(g.GoodStack, card)
		cardType = "GOOD"
		drinkType = "give"
	case 2: // Bad
		g.BadStack = append(g.BadStack, card)
		cardType = "BAD"
		drinkType = "take"
	case 0: // Ugly
		g.UglyStack = append(g.UglyStack, card)
		cardType = "UGLY"
		drinkType = "take"
	}
	
	// Check for matches
	matches := make([]PlayerMatch, 0)
	for _, player := range g.Players {
		for _, handCard := range player.Hand {
			if handCard.Rank == card.Rank {
				drinkSeconds := 1
				if cardType == "UGLY" {
					drinkSeconds = int(card.Rank)
				}
				
				matches = append(matches, PlayerMatch{
					PlayerID:     player.ID,
					HandCard:     handCard,
					DrinkSeconds: drinkSeconds,
					DrinkType:    drinkType,
				})
			}
		}
	}
	
	return &MainGameResult{
		Card:     card,
		CardType: cardType,
		Matches:  matches,
	}, nil
}

type PlayerMatch struct {
	PlayerID     int       `json:"playerId"`
	HandCard     deck.Card `json:"handCard"`
	DrinkSeconds int       `json:"drinkSeconds"`
	DrinkType    string    `json:"drinkType"` // "give" or "take"
}

type MainGameResult struct {
	Card     deck.Card      `json:"card"`
	CardType string         `json:"cardType"`
	Matches  []PlayerMatch  `json:"matches"`
}

func (g *Game) processBonus() (*MainGameResult, error) {
	if len(g.BonusUgly) == 0 {
		g.Phase = Finished
		return &MainGameResult{CardType: "FINISHED"}, nil
	}
	
	card := g.BonusUgly[0]
	g.BonusUgly = g.BonusUgly[1:]
	
	matches := make([]PlayerMatch, 0)
	for _, player := range g.Players {
		for _, handCard := range player.Hand {
			if handCard.Rank == card.Rank {
				matches = append(matches, PlayerMatch{
					PlayerID:     player.ID,
					HandCard:     handCard,
					DrinkSeconds: int(card.Rank),
					DrinkType:    "take",
				})
			}
		}
	}
	
	if len(g.BonusUgly) == 0 {
		g.Phase = Finished
	}
	
	return &MainGameResult{
		Card:     card,
		CardType: "BONUS UGLY",
		Matches:  matches,
	}, nil
}

// GetGameState returns the current game state
func (g *Game) GetGameState() map[string]interface{} {
	return map[string]interface{}{
		"id":            g.ID,
		"phase":         g.Phase.String(),
		"currentPlayer": g.CurrentPlayer,
		"round":         g.Round,
		"players":       g.Players,
		"cardsLeft":     len(g.Cards),
		"goodStack":     len(g.GoodStack),
		"badStack":      len(g.BadStack),
		"uglyStack":     len(g.UglyStack),
		"bonusUgly":     len(g.BonusUgly),
	}
}
