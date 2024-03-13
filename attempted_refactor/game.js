document.getElementById("startGame").addEventListener("click", function() {
    fetch('/start', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({playerCount: 2, deckCount: 1}) // Example setup
    })
    .then(response => response.json())
    .then(data => {
        console.log(data.message);
        document.getElementById("currentPlayer").textContent = "1";
        document.getElementById("drawCard").style.display = "block";
    });
});

document.getElementById("drawCard").addEventListener("click", function() {
    fetch('/draw', {method: 'POST'})
    .then(response => response.json())
    .then(data => {
        console.log('Card drawn:', data.card);
        const cardInfo = `Card: ${data.card.Value} of ${data.card.Suit}`;
        document.getElementById("cardDrawn").textContent = cardInfo;
        document.getElementById("cardDrawn").style.display = "block";
        // Implement changing turns or other game logic here
    });
});


document.getElementById("smoke").addEventListener("click", function() { makeChoice("Smoke"); });
document.getElementById("fire").addEventListener("click", function() { makeChoice("Fire"); });

function makeChoice(choice) {
    const currentPlayer = document.getElementById("currentPlayer").textContent;
    fetch('/makeChoice', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({playerNumber: currentPlayer, choice: choice})
    })
    .then(response => response.json())
    .then(data => {
        console.log(data.status);
        // Move to the next player or handle game logic as needed
        nextPlayer();
    });
}

function nextPlayer() {
    fetch('/next', {method: 'POST'})
    .then(response => response.json())
    .then(data => {
        document.getElementById("currentPlayer").textContent = data.currentPlayer;
        // Hide choices, prepare for the next player's turn, etc.
        document.getElementById("choices").style.display = "none";
    });
}
