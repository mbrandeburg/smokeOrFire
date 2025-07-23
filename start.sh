#!/bin/bash

# Add Go to PATH if it's not already there
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin

echo "🔥 Starting Smoke or Fire web server..."
echo "🎮 Game will be available at: http://localhost:8343"
echo "🃏 Ready to play!"

# Try to open browser automatically (works on most Linux/macOS systems)
sleep 2
if command -v xdg-open > /dev/null; then
    echo "🌐 Opening browser..."
    xdg-open http://localhost:8343 >/dev/null 2>&1 &
elif command -v open > /dev/null; then
    echo "🌐 Opening browser..."
    open http://localhost:8343 >/dev/null 2>&1 &
else
    echo "Please open your browser and go to: http://localhost:8343"
fi

# Start the server
go run server.go
