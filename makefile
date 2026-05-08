# --- Variables ---
DB_URL=postgres://admin:password@localhost:5432/costaBackend?sslmode=disable
MIGRATIONS_PATH=file://internal/db/migrations

# --- Commands ---
.PHONY: build dev run lint migrate-up migrate-down help

help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  build         Build the binary"
	@echo "  dev           Run the app with live reload"
	@echo "  run           Run the app directly"
	@echo "  lint          Run golangci-lint"
	@echo "  migrate-up    Run all up migrations"
	@echo "  migrate-down  Rollback the last migration"

build:
	go build -o bin/app ./cmd/api/main.go

dev:
	go run ./cmd/api/main.go

run:
	./bin/app

lint:
	golangci-lint run ./...

migrate-up:
	migrate -source $(MIGRATIONS_PATH) -database "$(DB_URL)" up

migrate-down:
	migrate -source $(MIGRATIONS_PATH) -database "$(DB_URL)" down 1