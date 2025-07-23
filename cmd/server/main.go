package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"smokeorfire/pkg/game"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all origins for development
		},
	}
	
	// Store active games and connections
	games       = make(map[string]*game.Game)
	connections = make(map[string][]*websocket.Conn)
	gamesMutex  = sync.RWMutex{}
)

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type CreateGameRequest struct {
	PlayerCount int `json:"playerCount"`
	DeckCount   int `json:"deckCount"`
}

type JoinGameRequest struct {
	GameID string `json:"gameId"`
}

type GuessRequest struct {
	GameID string         `json:"gameId"`
	Guess  game.GuessType `json:"guess"`
}

func main() {
	r := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	r.Use(cors.New(config))

	// Serve static files
	r.Static("/static", "./web/static")
	r.LoadHTMLGlob("web/templates/*")

	// Routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/api/games", createGame)
	r.GET("/api/games/:gameId", getGameState)
	r.POST("/api/games/:gameId/guess", makeGuess)
	r.POST("/api/games/:gameId/main-game", processMainGame)
	r.GET("/ws/:gameId", handleWebSocket)

	log.Println("Server starting on :8080")
	r.Run(":8080")
}

func createGame(c *gin.Context) {
	var req CreateGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gameID := generateGameID()
	newGame, err := game.NewGame(gameID, req.PlayerCount, req.DeckCount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gamesMutex.Lock()
	games[gameID] = newGame
	connections[gameID] = make([]*websocket.Conn, 0)
	gamesMutex.Unlock()

	c.JSON(http.StatusCreated, gin.H{
		"gameId": gameID,
		"game":   newGame.GetGameState(),
	})
}

func getGameState(c *gin.Context) {
	gameID := c.Param("gameId")
	
	gamesMutex.RLock()
	g, exists := games[gameID]
	gamesMutex.RUnlock()
	
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	c.JSON(http.StatusOK, g.GetGameState())
}

func makeGuess(c *gin.Context) {
	gameID := c.Param("gameId")
	
	var req GuessRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gamesMutex.Lock()
	g, exists := games[gameID]
	gamesMutex.Unlock()
	
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	// Start game if it's in setup phase
	if g.Phase == game.Setup {
		if err := g.StartGame(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	result, err := g.ProcessGuess(g.Players[g.CurrentPlayer].ID, req.Guess)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Broadcast update to all connected clients
	broadcastToGame(gameID, Message{
		Type: "gameUpdate",
		Data: map[string]interface{}{
			"gameState": g.GetGameState(),
			"result":    result,
		},
	})

	c.JSON(http.StatusOK, result)
}

func processMainGame(c *gin.Context) {
	gameID := c.Param("gameId")
	
	gamesMutex.Lock()
	g, exists := games[gameID]
	gamesMutex.Unlock()
	
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	result, err := g.ProcessMainGameCard()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Broadcast update to all connected clients
	broadcastToGame(gameID, Message{
		Type: "mainGameUpdate",
		Data: map[string]interface{}{
			"gameState": g.GetGameState(),
			"result":    result,
		},
	})

	c.JSON(http.StatusOK, result)
}

func handleWebSocket(c *gin.Context) {
	gameID := c.Param("gameId")
	
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	defer conn.Close()

	// Add connection to game
	gamesMutex.Lock()
	connections[gameID] = append(connections[gameID], conn)
	g, exists := games[gameID]
	gamesMutex.Unlock()

	if !exists {
		conn.WriteJSON(Message{Type: "error", Data: "Game not found"})
		return
	}

	// Send initial game state
	conn.WriteJSON(Message{
		Type: "gameState",
		Data: g.GetGameState(),
	})

	// Handle incoming messages
	for {
		var msg Message
		if err := conn.ReadJSON(&msg); err != nil {
			log.Printf("WebSocket read error: %v", err)
			break
		}

		// Process message based on type
		switch msg.Type {
		case "ping":
			conn.WriteJSON(Message{Type: "pong", Data: nil})
		}
	}

	// Remove connection when done
	gamesMutex.Lock()
	conns := connections[gameID]
	for i, c := range conns {
		if c == conn {
			connections[gameID] = append(conns[:i], conns[i+1:]...)
			break
		}
	}
	gamesMutex.Unlock()
}

func broadcastToGame(gameID string, message Message) {
	gamesMutex.RLock()
	conns := connections[gameID]
	gamesMutex.RUnlock()

	for _, conn := range conns {
		if err := conn.WriteJSON(message); err != nil {
			log.Printf("Failed to send message to WebSocket: %v", err)
		}
	}
}

func generateGameID() string {
	// Simple game ID generation - you might want something more sophisticated
	return strconv.FormatInt(1000000+int64(len(games)), 10)
}
