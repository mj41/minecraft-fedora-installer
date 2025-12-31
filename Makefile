.PHONY: build clean install

# Binary name
BINARY=mc-installer

# Build the installer
build:
	go build -o $(BINARY) main.go

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

# Run with verbose output
run: build
	./$(BINARY)

# Run with force flag
run-force: build
	./$(BINARY) --force
