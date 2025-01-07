include .env

.PHONY: build lint migrate-create migrate-up migrate-down

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

# Perform migrations file creation
migrate-create:
	migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

# Perform up migrations
migrate-up:
	migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDRESS) up

# Perform down migrations
migrate-down:
	migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDRESS) down
