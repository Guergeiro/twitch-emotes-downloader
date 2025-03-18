# Simple Makefile for a Go project
BUILD_FOLDER?="build"
GOARCH?=amd64
GOOS?=linux

all: build

# Build the application
build:
ifeq ($(GOOS), windows)
build: build-windows
else
build: build-unix
endif

build-unix:
	@echo "Building Unix..."
	@CGO_ENABLED=0 go build -o "$(BUILD_FOLDER)/twe-dl-$(GOOS)-$(GOARCH)" cmd/cli/main.go

build-windows:
	@echo "Building Windows..."
	@CGO_ENABLED=0 go build -o "$(BUILD_FOLDER)/twe-dl-$(GOOS)-$(GOARCH).exe" cmd/cli/main.go

archive:
ifeq ($(GOOS), windows)
archive: archive-windows
else
archive: archive-unix
endif

archive-unix:
	@echo "Archiving Unix..."
	@tar czvf "twe-dl-$(GOOS)-$(GOARCH).tar.gz" "$(BUILD_FOLDER)/twe-dl-$(GOOS)-$(GOARCH)"

archive-windows:
	@echo "Archiving Windows..."
	@7z a "twe-dl-$(GOOS)-$(GOARCH).zip" "$(BUILD_FOLDER)/twe-dl-$(GOOS)-$(GOARCH).exe"

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

.PHONY: all build build-unix build-windows archive archive-unix archive-windows run test clean change-add change-empty change-status change-version change-tag
