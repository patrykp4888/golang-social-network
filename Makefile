.PHONY: build lint

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=golang-social-network
LINTER=golangci-lint

# Build the project
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

# Lint the project
lint:
	$(LINTER) run
