.PHONY: build

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=golang-social-network

# Build the project
build:
	$(GOBUILD) -o $(BINARY_NAME) -v


