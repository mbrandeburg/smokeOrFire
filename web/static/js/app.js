class SmokeOrFireApp {
    constructor() {
        this.gameId = null;
        this.gameState = null;
        this.websocket = null;
        this.currentScreen = 'welcome';
        
        this.initializeEventListeners();
        // Remove the automatic welcome sequence - we'll trigger it when game starts
    }

    initializeEventListeners() {
        // Welcome screen
        document.getElementById('create-game-btn').addEventListener('click', () => this.createGame());
        document.getElementById('join-game-btn').addEventListener('click', () => this.joinGame());
        
        // Player count change handler
        document.getElementById('player-count').addEventListener('change', (e) => this.updatePlayerNameInputs(e.target.value));
        
        // Main game
        document.getElementById('draw-card-btn').addEventListener('click', () => this.drawMainGameCard());
        
        // Enter key handlers
        document.getElementById('game-id').addEventListener('keypress', (e) => {
            if (e.key === 'Enter') this.joinGame();
        });

        // Initialize player name inputs
        this.updatePlayerNameInputs(4);
    }

    updatePlayerNameInputs(playerCount) {
        const container = document.getElementById('player-names-inputs');
        container.innerHTML = '';
        
        for (let i = 1; i <= playerCount; i++) {
            const input = document.createElement('input');
            input.type = 'text';
            input.id = `player-name-${i}`;
            input.placeholder = `Player ${i} name (optional)`;
            input.className = 'player-name-input';
            container.appendChild(input);
        }
    }

    async createGame() {
        const playerCount = parseInt(document.getElementById('player-count').value);
        const deckCount = parseInt(document.getElementById('deck-count').value);

        if (playerCount < 2 || playerCount > 20) {
            alert('Player count must be between 2 and 20');
            return;
        }

        if (deckCount < 1 || deckCount > 20) {
            alert('Deck count must be between 1 and 20');
            return;
        }

        // Collect player names
        const playerNames = [];
        for (let i = 1; i <= playerCount; i++) {
            const nameInput = document.getElementById(`player-name-${i}`);
            playerNames.push(nameInput ? nameInput.value.trim() : '');
        }

        this.showLoading(true);

        try {
            const response = await fetch('/api/games', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    playerCount,
                    deckCount,
                    playerNames
                })
            });

            const data = await response.json();
            
            if (response.ok) {
                this.gameId = data.gameId;
                this.gameState = data.game;
                this.connectWebSocket();
                this.showGameScreen();
            } else {
                alert('Error creating game: ' + data.error);
            }
        } catch (error) {
            alert('Error creating game: ' + error.message);
        } finally {
            this.showLoading(false);
        }
    }

    async joinGame() {
        const gameId = document.getElementById('game-id').value.trim();
        
        if (!gameId) {
            alert('Please enter a game ID');
            return;
        }

        this.showLoading(true);

        try {
            const response = await fetch(`/api/games/${gameId}`);
            const data = await response.json();
            
            if (response.ok) {
                this.gameId = gameId;
                this.gameState = data;
                this.connectWebSocket();
                this.showGameScreen();
            } else {
                alert('Error joining game: ' + data.error);
            }
        } catch (error) {
            alert('Error joining game: ' + error.message);
        } finally {
            this.showLoading(false);
        }
    }

    connectWebSocket() {
        const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
        const wsUrl = `${protocol}//${window.location.host}/ws/${this.gameId}`;
        
        this.websocket = new WebSocket(wsUrl);
        
        this.websocket.onopen = () => {
            console.log('WebSocket connected');
        };
        
        this.websocket.onmessage = (event) => {
            const message = JSON.parse(event.data);
            this.handleWebSocketMessage(message);
        };
        
        this.websocket.onclose = () => {
            console.log('WebSocket disconnected');
            // Attempt to reconnect after 3 seconds
            setTimeout(() => {
                if (this.gameId) {
                    this.connectWebSocket();
                }
            }, 3000);
        };
        
        this.websocket.onerror = (error) => {
            console.error('WebSocket error:', error);
        };
    }

    handleWebSocketMessage(message) {
        switch (message.type) {
            case 'gameState':
                this.gameState = message.data;
                this.updateGameDisplay();
                break;
            case 'gameUpdate':
                this.gameState = message.data.gameState;
                this.updateGameDisplay();
                this.showResult(message.data.result);
                break;
            case 'mainGameUpdate':
                this.gameState = message.data.gameState;
                this.updateGameDisplay();
                this.showMainGameResult(message.data.result);
                break;
            case 'error':
                alert('Game error: ' + message.data);
                break;
        }
    }

    showGameScreen() {
        document.getElementById('welcome-screen').classList.remove('active');
        document.getElementById('game-screen').classList.add('active');
        this.currentScreen = 'game';
        
        // Show the classic "I have three questions" message when game starts
        this.showThreeQuestionsMessage();
        
        this.updateGameDisplay();
    }

    showThreeQuestionsMessage() {
        // Create the message overlay
        const overlay = document.createElement('div');
        overlay.className = 'three-questions-overlay';
        overlay.innerHTML = `
            <div class="three-questions-content">
                <h2 style="font-size: 2rem; margin-bottom: 20px; animation: fadeInUp 1s ease-out;">I have three questions for you...</h2>
                <p style="animation: fadeInUp 1s ease-out 1s both;">Let's begin with the first question:</p>
            </div>
        `;
        
        document.body.appendChild(overlay);
        
        // Remove the overlay after 4 seconds
        setTimeout(() => {
            overlay.style.opacity = '0';
            setTimeout(() => {
                if (overlay.parentNode) {
                    overlay.parentNode.removeChild(overlay);
                }
            }, 500);
        }, 4000);
    }

    updateGameDisplay() {
        if (!this.gameState) return;

        // Update header info
        document.getElementById('current-game-id').textContent = this.gameId;
        document.getElementById('current-phase').textContent = this.gameState.phase;
        document.getElementById('current-round').textContent = this.gameState.round;
        document.getElementById('cards-left').textContent = this.gameState.cardsLeft;

        // Update players
        this.updatePlayersDisplay();

        // Update game controls based on phase
        this.updateGameControls();

        // Update stacks info for main game
        if (this.gameState.phase === 'MainGame') {
            document.getElementById('good-count').textContent = this.gameState.goodStack;
            document.getElementById('bad-count').textContent = this.gameState.badStack;
            document.getElementById('ugly-count').textContent = this.gameState.uglyStack;
        }
    }

    updatePlayersDisplay() {
        const playersContainer = document.getElementById('players-list');
        playersContainer.innerHTML = '';

        this.gameState.players.forEach((player, index) => {
            const playerDiv = document.createElement('div');
            playerDiv.className = `player-card ${index === this.gameState.currentPlayer ? 'active' : ''}`;
            
            const handHTML = player.hand.map(card => {
                const isRed = card.suit === 1 || card.suit === 3; // Diamond or Heart
                return `<span class="card ${isRed ? 'red' : 'black'}">${this.formatCard(card)}</span>`;
            }).join('');

            playerDiv.innerHTML = `
                <div class="player-name">${player.name}</div>
                <div class="player-hand">${handHTML}</div>
            `;
            
            playersContainer.appendChild(playerDiv);
        });
    }

    updateGameControls() {
        const currentPlayerSection = document.getElementById('current-player-section');
        const mainGameSection = document.getElementById('main-game-section');
        const guessButtonsContainer = document.getElementById('guess-buttons');
        
        if (this.gameState.phase === 'MainGame') {
            currentPlayerSection.classList.add('hidden');
            mainGameSection.classList.remove('hidden');
            return;
        }

        if (this.gameState.phase === 'Finished') {
            currentPlayerSection.innerHTML = '<h3>Game Finished!</h3><p>Thanks for playing!</p>';
            mainGameSection.classList.add('hidden');
            return;
        }

        // Show current player section
        currentPlayerSection.classList.remove('hidden');
        mainGameSection.classList.add('hidden');

        // Update current player name
        const currentPlayer = this.gameState.players[this.gameState.currentPlayer];
        document.getElementById('current-player-name').textContent = `${currentPlayer.name}'s Turn`;

        // Update instruction and buttons based on phase
        const instruction = document.getElementById('game-instruction');
        guessButtonsContainer.innerHTML = '';

        let buttons = [];
        switch (this.gameState.phase) {
            case 'SmokeOrFire':
                instruction.textContent = 'Choose Smoke or Fire:';
                buttons = [
                    { text: 'Smoke', value: 'smoke', class: 'btn-primary' },
                    { text: 'Fire', value: 'fire', class: 'btn-secondary' }
                ];
                break;
            case 'HigherOrLower':
                instruction.textContent = 'Will the next card be Higher or Lower?';
                buttons = [
                    { text: 'Higher', value: 'higher', class: 'btn-primary' },
                    { text: 'Lower', value: 'lower', class: 'btn-secondary' }
                ];
                break;
            case 'InsideOrOutside':
                instruction.textContent = 'Will the next card be Inside or Outside your range?';
                buttons = [
                    { text: 'Inside', value: 'inside', class: 'btn-primary' },
                    { text: 'Outside', value: 'outside', class: 'btn-secondary' }
                ];
                break;
            case 'OddOrEven':
                instruction.textContent = 'Will the next card be Odd or Even?';
                buttons = [
                    { text: 'Odd', value: 'odd', class: 'btn-primary' },
                    { text: 'Even', value: 'even', class: 'btn-secondary' }
                ];
                break;
        }

        buttons.forEach(button => {
            const btn = document.createElement('button');
            btn.className = `btn ${button.class} guess-btn`;
            btn.textContent = button.text;
            
            btn.addEventListener('click', () => this.makeGuess(button.value));
            
            guessButtonsContainer.appendChild(btn);
        });
    }

    async makeGuess(guess) {
        try {
            const response = await fetch(`/api/games/${this.gameId}/guess`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ guess })
            });

            const data = await response.json();
            
            if (!response.ok) {
                alert('Error making guess: ' + data.error);
            }
        } catch (error) {
            alert('Error making guess: ' + error.message);
        }
    }

    async drawMainGameCard() {
        try {
            const response = await fetch(`/api/games/${this.gameId}/main-game`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                }
            });

            const data = await response.json();
            
            if (!response.ok) {
                alert('Error drawing card: ' + data.error);
            }
        } catch (error) {
            alert('Error drawing card: ' + error.message);
        }
    }

    showResult(result) {
        const resultsContainer = document.getElementById('results-list');
        
        const resultDiv = document.createElement('div');
        resultDiv.className = `result-item ${result.isCorrect ? 'correct' : 'incorrect'}`;
        
        let drinkText = '';
        if (result.drinkSeconds > 0) {
            drinkText = ` - Drink for ${result.drinkSeconds} second${result.drinkSeconds > 1 ? 's' : ''}!`;
        }
        
        resultDiv.innerHTML = `
            <strong>Player ${result.playerId}</strong> drew ${this.formatCard(result.card)} 
            and guessed ${result.guess}. ${result.isCorrect ? 'Correct!' : 'Wrong!'}${drinkText}
        `;
        
        resultsContainer.insertBefore(resultDiv, resultsContainer.firstChild);
        
        // Keep only last 5 results
        while (resultsContainer.children.length > 5) {
            resultsContainer.removeChild(resultsContainer.lastChild);
        }
    }

    showMainGameResult(result) {
        if (result.cardType === 'FINISHED') {
            return;
        }

        const resultsContainer = document.getElementById('results-list');
        
        const resultDiv = document.createElement('div');
        resultDiv.className = 'result-item';
        
        let matchesText = '';
        if (result.matches && result.matches.length > 0) {
            const matchList = result.matches.map(match => {
                const action = match.drinkType === 'give' ? 'gives a drink' : `drinks for ${match.drinkSeconds} second${match.drinkSeconds > 1 ? 's' : ''}`;
                return `Player ${match.playerId} ${action}`;
            }).join(', ');
            matchesText = ` - Matches: ${matchList}`;
        } else {
            matchesText = ' - No matches';
        }
        
        resultDiv.innerHTML = `
            <strong>${result.cardType}</strong>: ${this.formatCard(result.card)}${matchesText}
        `;
        
        resultsContainer.insertBefore(resultDiv, resultsContainer.firstChild);
        
        // Keep only last 5 results
        while (resultsContainer.children.length > 5) {
            resultsContainer.removeChild(resultsContainer.lastChild);
        }
    }

    formatCard(card) {
        const suits = ['♠', '♦', '♣', '♥'];
        const ranks = ['', 'A', '2', '3', '4', '5', '6', '7', '8', '9', '10', 'J', 'Q', 'K'];
        
        return `${ranks[card.rank]}${suits[card.suit]}`;
    }

    showLoading(show) {
        const loading = document.getElementById('loading');
        if (show) {
            loading.classList.remove('hidden');
        } else {
            loading.classList.add('hidden');
        }
    }
}

// Initialize the app when the page loads
document.addEventListener('DOMContentLoaded', () => {
    new SmokeOrFireApp();
});
