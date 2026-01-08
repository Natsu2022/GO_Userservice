# =============================================================
#                   User Service Makefile
# =============================================================

APP_NAME=user_service
CMD_DIR=../cmd/$(APP_NAME)
OUTPUT=../../$(APP_NAME)/$(APP_NAME).exe

# Detect OS
OS := $(shell uname -s)

# =============================================================
#                  Environment Variables
# =============================================================

export APP_ENV ?= dev   # default = dev

# =============================================================
#                  Development Commands
# =============================================================

dev:
	@echo "🚀 Building $(APP_NAME) in DEV mode..."
	@APP_ENV=prod go build -ldflags="-s -w" -o $(OUTPUT) $(CMD_DIR)/main.go
	@echo "▶ Running $(APP_NAME)..."
	@APP_ENV=dev ./$(OUTPUT)

run:
	@echo "▶ Running service directly..."
	@APP_ENV=dev ./$(OUTPUT)

watch:
	@echo "👀 Watching for code changes..."
	@reflex -r '\.go$$' -- sh -c "APP_ENV=dev go run $(CMD_DIR)/main.go"

# =============================================================
#                  Build Commands (Production)
# =============================================================

build:
	@echo "🏗️ Building $(APP_NAME) for production..."
	@APP_ENV=prod go build -ldflags="-s -w" -o $(OUTPUT) $(CMD_DIR)/main.go
	@echo "✅ Production binary created → $(OUTPUT)"

release:
	@echo "📦 Building Linux release..."
	@GOOS=linux GOARCH=amd64 APP_ENV=prod go build -ldflags="-s -w" -o $(APP_NAME)
	@echo "👍 Linux binary created: $(APP_NAME)"

# =============================================================
#                       Docker Commands
# =============================================================

docker-build:
	@echo "🐳 Building Docker image..."
	@docker build -t $(APP_NAME):latest ../../
	@echo "✅ Docker image built."

docker-run:
	@echo "🚢 Running Docker container..."
	@docker run -p 3000:3000 $(APP_NAME):latest

# =============================================================
#                       Tools / Clean
# =============================================================

tidy:
	@echo "🧹 Cleaning unused Go modules..."
	@go mod tidy

fmt:
	@echo "🪄 Formatting Go code..."
	@go fmt ./...

clean:
	@echo "🗑️ Cleaning build files..."
	@rm -f $(APP_NAME).exe
	@rm -f $(OUTPUT)
	@echo "✨ Cleanup done."

# =============================================================
#                       Migration
# =============================================================
DB_URL=postgres://root:root1234@localhost:5434/mydb?sslmode=disable

migrate-up:
	@migrate -path ../migrations -database "$(DB_URL)" up

migrate-drop:
	@migrate -path ../migrations -database "$(DB_URL)" drop -f

# =============================================================
#                       Help (default)
# =============================================================
help:
	@echo ""
	@echo "Usage: make <command>"
	@echo ""
	@echo "Commands:"
	@echo "  dev           Build + run service in DEV mode"
	@echo "  run           Run existing dev build"
	@echo "  watch         Auto-restart on changes (requires reflex)"
	@echo ""
	@echo "  build         Production build (Windows)"
	@echo "  release       Production build (Linux)"
	@echo ""
	@echo "  docker-build  Build Docker image"
	@echo "  docker-run    Run Docker container"
	@echo ""
	@echo "  tidy          Clean Go modules"
	@echo "  fmt           Format Go code"
	@echo "  clean         Remove build artifacts"
	@echo ""
	@echo "  migrate-up	   Setup Database"
	@echo "  migrate-Drop	   Drop All Database"