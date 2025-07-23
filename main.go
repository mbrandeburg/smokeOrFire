// Smoke Or Fire - Web Application
// The card game introduced by Kyle Durham. 
// Originally coded by Matthew Brandeburg in January 2019
// Refactored to web application in 2025

package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	log.Println("🔥 Starting Smoke or Fire web application...")
	
	// Try to open the browser automatically
	go func() {
		// Wait a moment for the server to start
		time.Sleep(2 * time.Second)
		
		url := "http://localhost:8080"
		var err error
		
		switch runtime.GOOS {
		case "linux":
			err = exec.Command("xdg-open", url).Start()
		case "windows":
			err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		case "darwin":
			err = exec.Command("open", url).Start()
		}
		
		if err != nil {
			log.Printf("Could not open browser automatically: %v", err)
		}
		log.Printf("🎮 Game available at: %s", url)
	}()
	
	// Run the server directly
	cmd := exec.Command("go", "run", "server.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = "."
	
	if err := cmd.Run(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
