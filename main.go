package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gophercises/smokeOrFire/deck"
)

// Player represents a player in the game
type Player struct {
	Number int
	Hand   []deck.Card
}

// Game represents the game state
type Game struct {
	Players   []*Player
	DeckCount int
	Cards     []deck.Card
}

// newGame creates a new game with the given number of players and decks
func newGame(playerCount, deckCount int) *Game {
	players := make([]*Player, playerCount)
	for i := range players {
		players[i] = &Player{Number: i + 1}
	}

	cards := deck.New(deck.Deck(deckCount), deck.Shuffle)
	return &Game{Players: players, DeckCount: deckCount, Cards: cards}
}

// drawCard draws a card from the game's deck
func (g *Game) drawCard() (deck.Card, []deck.Card) {
	return g.Cards[0], g.Cards[1:]
}

// renderTemplate renders the given template with the provided data
func renderTemplate(c *gin.Context, tmplStr string, data interface{}) {
	tmpl, err := template.New("").Parse(tmplStr)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error parsing template: %v", err))
		return
	}
	if tmpl == nil {
		c.String(http.StatusInternalServerError, "Failed to create template")
		return
	}
	c.HTML(http.StatusOK, tmpl.Name(), data)
}

func main() {
	r := gin.Default()

	// Render the game setup form
	r.GET("/", func(c *gin.Context) {
		renderTemplate(c, setupTemplate, nil)
	})

	// Handle game setup form submission
	r.POST("/setup", func(c *gin.Context) {
		playerCount, deckCount := parseSetupForm(c)
		game := newGame(playerCount, deckCount)
		c.Redirect(http.StatusFound, "/game")
		// Store the game state in the session or other storage
		storeGameInStorage(c, game)
	})

	// Render the game page
	r.GET("/game", func(c *gin.Context) {
		// Fetch the game state from the session or other storage
		game := getGameFromStorage(c)
		renderTemplate(c, gameTemplate, game)
	})

	// Handle game actions (e.g., make a guess, draw a card)
	r.POST("/game", func(c *gin.Context) {
		// Fetch the game state from the session or other storage
		game := getGameFromStorage(c)

		// Process the user's action and update the game state
		processAction(game, c.PostForm("action"))

		// Store the updated game state
		storeGameInStorage(c, game)

		// Redirect back to the game page
		c.Redirect(http.StatusFound, "/game")
	})

	r.Run(":8080")
}

// parseSetupForm parses the player count and deck count from the setup form
func parseSetupForm(c *gin.Context) (int, int) {
	playerCount := parseInt(c.PostForm("playerCount"))
	deckCount := parseInt(c.PostForm("deckCount"))
	return playerCount, deckCount
}

// parseInt converts a string to an integer, with default value 0
func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// processAction processes the user's action and updates the game state
func processAction(game *Game, action string) {
	// Implement your game logic here
	// For example:
	if action == "drawCard" {
		player := game.Players[0] // Assuming a single-player game for simplicity
		card, remaining := game.drawCard()
		player.Hand = append(player.Hand, card)
		game.Cards = remaining
	}
}

// getGameFromStorage retrieves the game state from the session or other storage
func getGameFromStorage(c *gin.Context) *Game {
	// Implement your storage logic here
	// For example, if you're storing the game state in the session:
	// return c.MustGet("game").(*Game)
	return nil
}

// storeGameInStorage stores the game state in the session or other storage
func storeGameInStorage(c *gin.Context, game *Game) {
	// Implement your storage logic here
	// For example, if you're storing the game state in the session:
	// c.Set("game", game)
}

// setupTemplate is the HTML template for the game setup form
const setupTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Smoke or Fire</title>
</head>
<body>
    <h1>Smoke or Fire</h1>
    <form method="post" action="/setup">
        <label>Number of Players:</label>
        <input type="number" name="playerCount" min="2" max="20" required>
        <br>
        <label>Number of Decks:</label>
        <input type="number" name="deckCount" min="1" max="20" required>
        <br>
        <button type="submit">Start Game</button>
    </form>
</body>
</html>
`

// gameTemplate is the HTML template for the game page
const gameTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Smoke or Fire</title>
</head>
<body>
    <h1>Smoke or Fire</h1>
    {{ range .Players }}
        <p>Player {{ .Number }}: {{ range .Hand }}{{ . }} {{ end }}</p>
    {{ end }}
    <form method="post" action="/game">
        <button type="submit" name="action" value="drawCard">Draw Card</button>
    </form>
</body>
</html>
`
