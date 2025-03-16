# Simple Makefile for a Go project
BUILD_FOLDER?="build"
GOARCH?=amd64

all: build

# Build the application

build: build-unix build-windows

build-unix:
	@echo "Building Unix..."
	@CGO_ENABLED=0 go build -o "$(BUILD_FOLDER)/twe-dl-$(GOARCH)" cmd/cli/main.go

build-windows:
	@echo "Building Windows..."
	@CGO_ENABLED=0 go build -o "$(BUILD_FOLDER)/twe-dl-$(GOARCH).exe" cmd/cli/main.go

archive: archive-linux archive-macos archive-windows

archive-linux:
	@echo "Archiving Linux..."
	@tar czvf "twe-dl-linux-$(GOARCH).tar.gz" "$(BUILD_FOLDER)/twe-dl-$(GOARCH)"

archive-macos:
	@echo "Archiving MacOS..."
	@tar czvf "twe-dl-macos-$(GOARCH).tar.gz" "$(BUILD_FOLDER)/twe-dl-$(GOARCH)"

archive-windows:
	@echo "Archiving Windows..."
	@7z a "twe-dl-windows-$(GOARCH).zip" "$(BUILD_FOLDER)/twe-dl-$(GOARCH).exe"

# Run the application
run:
	@go run cmd/cli/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./...

coverage:
	@echo "Testing..."
	@go test ./... -coverprofile="c.out"
	@go tool cover -html="c.out"

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -rfd build

# Changesets tasks
change-add:
	@pnpm dlx @changesets/cli add

change-chore:
	@pnpm dlx @changesets/cli add --empty

change-status:
	@pnpm dlx @changesets/cli status --since=origin/master

change-version:
	@pnpm dlx @changesets/cli version

change-tag:
	@pnpm dlx @changesets/cli tag

.PHONY: all build run test clean change-add change-empty change-status change-version change-tag
