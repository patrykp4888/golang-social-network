include .env

.PHONY: build lint migrate-up migrate-down

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=golang-social-network
LINTER=golangci-lint

# Migrations parameters
MIGRATIONS_PATH=./cmd/migrate/migrations

# Build the project
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

# Lint the project
lint:
	$(LINTER) run

# Perform up migrations
migrate-up:
	migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDRESS) up

# Perform down migrations
migrate-down:
	migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDRESS) down
