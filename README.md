# Smoke or Fire - Web Application

A modern web-based version of the card game "Smoke or Fire" introduced by Kyle Durham, originally coded by Matthew Brandeburg in January 2019 and refactored to a web application in 2025.

## 🎮 About the Game

Smoke or Fire is a drinking card game with multiple phases:

1. **Smoke or Fire** - Players guess if their card will be black (smoke) or red (fire)
2. **Higher or Lower** - Players guess if the next card will be higher or lower than their previous card
3. **Inside or Outside** - Players guess if their card will fall inside or outside the range of their first two cards
4. **Odd or Even** - Players guess if their card will be odd or even
5. **Main Game** - The "Good, Bad, and Ugly" phase where remaining cards are dealt and players match their hand cards

## 🚀 Features

- **Modern Web Interface** - Beautiful, responsive design that works on desktop and mobile
- **Real-time Multiplayer** - Multiple players can join the same game using WebSockets
- **Custom Player Names** - Players can set their own names for a personalized experience
- **Animated UI** - Smooth animations and visual feedback including the classic ASCII art
- **Game State Management** - Persistent game state across sessions
- **Cross-platform** - Works on Windows, macOS, and Linux
- **Docker Support** - Easy containerized deployment

## 🛠️ Installation & Setup

### Prerequisites
- Go 1.21 or later
- Modern web browser

### Quick Start

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd smokeOrFire
   ```

2. **Initialize Go modules**
   ```bash
   go mod tidy
   ```

3. **Run the application**
   
   **Windows:**
   ```powershell
   .\start.ps1
   ```
   
   **Linux/macOS:**
   ```bash
   chmod +x start.sh
   ./start.sh
   ```
   
   **Or manually:**
   ```bash
   go run server.go
   ```

4. **Open your browser**
   The application will be available at `http://localhost:8343`

### Docker Deployment

**Using Docker Compose (Recommended):**
```bash
docker-compose up -d
```

**Using Makefile:**
```bash
make docker-run
```

**Manual Docker:**
```bash
docker build -t smokeorfire .
docker run -p 8343:8343 smokeorfire
```

### Manual Server Start

If you prefer to run the server directly:

```bash
go run server.go
```

## 🎯 How to Play

1. **Create a Game**: Set the number of players (2-20), decks (1-20), and optionally player names
2. **Share Game ID**: Other players can join using the generated Game ID
3. **Play Through Phases**: Follow the on-screen instructions for each phase
4. **Main Game**: Watch as cards are dealt and see who gets matches!

## 🔧 Development Commands

```bash
# Development
make dev          # Run development server
make build        # Build binary

# Docker
make docker-build # Build Docker image  
make docker-run   # Run with Docker Compose
make docker-stop  # Stop containers
make docker-logs  # View logs
make clean        # Cleanup Docker resources
```

## 🏗️ Architecture

The application is built with:

- **Backend**: Go with Gin web framework
- **Frontend**: Vanilla JavaScript with modern CSS
- **Communication**: WebSockets for real-time updates
- **Game Logic**: Modular design with separate packages

### Project Structure

```
smokeOrFire/
├── main.go              # Application entry point
├── server.go            # Web server implementation
├── go.mod              # Go module dependencies
├── pkg/
│   ├── deck/           # Card deck implementation
│   └── game/           # Game logic and state management
└── web/
    ├── templates/      # HTML templates
    └── static/
        ├── css/        # Stylesheets
        └── js/         # JavaScript application
```

## 🔧 Development

### Adding New Features

1. **Game Logic**: Modify `pkg/game/game.go`
2. **API Endpoints**: Add routes in `server.go`
3. **Frontend**: Update `web/static/js/app.js`
4. **Styling**: Modify `web/static/css/style.css`

### API Endpoints

- `POST /api/games` - Create a new game
- `GET /api/games/:gameId` - Get game state
- `POST /api/games/:gameId/guess` - Make a guess
- `POST /api/games/:gameId/main-game` - Draw a card in main game
- `GET /ws/:gameId` - WebSocket connection for real-time updates

## 📱 Mobile Support

The application is fully responsive and works great on mobile devices. The touch-friendly interface makes it perfect for party games!

## 🤝 Contributing

Feel free to submit issues and enhancement requests!

## 📄 License

See LICENSE file for details.

## 🎉 Credits

- **Original Game**: Kyle Durham
- **Original Code**: Matthew Brandeburg (January 2019)
- **Web Refactor**: Matthew Brandeburg + Claude AI Agents (2025)

---

**Have fun and drink responsibly!** 🍻
