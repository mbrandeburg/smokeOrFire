$env:PATH += ";C:\Program Files\Go\bin"
Write-Host "🔥 Starting Smoke or Fire web server..." -ForegroundColor Red
Write-Host "🎮 Game will be available at: http://localhost:8343" -ForegroundColor Green  
Write-Host "🃏 Ready to play!" -ForegroundColor Yellow

# Try to open browser automatically
Start-Sleep -Seconds 2
try {
    Start-Process "http://localhost:8343"
    Write-Host "🌐 Opening browser..." -ForegroundColor Cyan
} catch {
    Write-Host "Could not open browser automatically. Please go to http://localhost:8343" -ForegroundColor Yellow
}

go run server.go
