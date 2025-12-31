.PHONY: build clean install

# Binary name
BINARY=mc-installer

# Version info
VERSION ?= dev
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS := -ldflags "-X main.Version=$(VERSION) -X main.GitCommit=$(GIT_COMMIT) -X main.BuildDate=$(BUILD_DATE)"

# Build the installer
build:
	go build $(LDFLAGS) -o $(BINARY) main.go

# Install (run the installer)
install: build
	./$(BINARY)

# Clean build artifacts
clean:
	rm -f $(BINARY)
	rm -f *.tar.gz

# Format code
fmt:
	go fmt ./...

# Validate desktop file (requires desktop-file-utils package)
validate-desktop:
	desktop-file-validate minecraft.desktop.tmpl

# Run with verbose output
run: build
	./$(BINARY)

# Run with force flag
run-force: build
	./$(BINARY) --force
