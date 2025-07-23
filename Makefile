# Makefile for Smoke or Fire

.PHONY: help build run stop clean docker-build docker-run docker-stop docker-logs dev

# Default target
help:
	@echo "🔥 Smoke or Fire - Available commands:"
	@echo ""
	@echo "Development:"
	@echo "  dev          - Run development server locally"
	@echo "  build        - Build the Go binary"
	@echo ""
	@echo "Docker:"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run with Docker Compose"
	@echo "  docker-stop  - Stop Docker containers"
	@echo "  docker-logs  - View Docker logs"
	@echo "  clean        - Clean up Docker images and containers"
	@echo ""

# Development
dev:
	@echo "🔥 Starting development server..."
	go run server.go

build:
	@echo "🔨 Building binary..."
	go build -o smokeorfire server.go

# Docker operations
docker-build:
	@echo "🐳 Building Docker image..."
	docker build -t smokeorfire:latest .

docker-run:
	@echo "🚀 Starting with Docker Compose..."
	docker-compose up -d
	@echo "🎮 Game available at: http://localhost:8343"

docker-stop:
	@echo "🛑 Stopping Docker containers..."
	docker-compose down

docker-logs:
	@echo "📋 Viewing Docker logs..."
	docker-compose logs -f

clean:
	@echo "🧹 Cleaning up..."
	docker-compose down
	docker system prune -f
	@echo "✅ Cleanup complete"

# Quick start with Docker
quick-start: docker-build docker-run
