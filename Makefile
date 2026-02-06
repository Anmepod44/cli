.PHONY: all build install uninstall clean test run setup deb snap help

# Variables
BINARY_NAME=cli-assistant
INSTALL_PATH=/usr/local/bin
VERSION=1.0.0

# Default target
all: build

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) ./cmd/cli-assistant
	@echo "✓ Build complete: ./$(BINARY_NAME)"

# Install to system path
install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_PATH)..."
	@sudo cp $(BINARY_NAME) $(INSTALL_PATH)/
	@sudo chmod +x $(INSTALL_PATH)/$(BINARY_NAME)
	@echo "✓ Installed to $(INSTALL_PATH)/$(BINARY_NAME)"
	@echo ""
	@echo "Run 'cli-assistant' to start the application"

# Uninstall from system
uninstall:
	@echo "Uninstalling $(BINARY_NAME)..."
	@sudo rm -f $(INSTALL_PATH)/$(BINARY_NAME)
	@echo "✓ Uninstalled"
	@echo ""
	@echo "Note: User data in ~/.cli-assistant is preserved"
	@echo "To remove: rm -rf ~/.cli-assistant"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -f $(BINARY_NAME)
	@rm -rf build/
	@rm -f *.deb
	@rm -f *.snap
	@go clean
	@echo "✓ Clean complete"

# Run tests
test:
	@echo "Running tests..."
	@go test ./... -v

# Run the application
run: build
	@./$(BINARY_NAME)

# Run setup script
setup:
	@./setup.sh

# Build Debian package
deb:
	@echo "Building Debian package..."
	@./build-deb.sh
	@echo "✓ Debian package created"

# Build Snap package
snap:
	@echo "Building Snap package..."
	@snapcraft
	@echo "✓ Snap package created"

# Development build with race detector
dev:
	@echo "Building with race detector..."
	@go build -race -o $(BINARY_NAME) ./cmd/cli-assistant
	@echo "✓ Development build complete"

# Build for multiple architectures
build-all:
	@echo "Building for multiple architectures..."
	@GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux-amd64 ./cmd/cli-assistant
	@GOOS=linux GOARCH=arm64 go build -o $(BINARY_NAME)-linux-arm64 ./cmd/cli-assistant
	@GOOS=linux GOARCH=386 go build -o $(BINARY_NAME)-linux-386 ./cmd/cli-assistant
	@echo "✓ Built for multiple architectures"

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "✓ Code formatted"

# Lint code
lint:
	@echo "Linting code..."
	@go vet ./...
	@echo "✓ Lint complete"

# Update dependencies
deps:
	@echo "Updating dependencies..."
	@go mod tidy
	@go mod download
	@echo "✓ Dependencies updated"

# Show help
help:
	@echo "CLI Command Assistant - Makefile targets:"
	@echo ""
	@echo "  make build       - Build the application"
	@echo "  make install     - Install to system path (requires sudo)"
	@echo "  make uninstall   - Remove from system"
	@echo "  make run         - Build and run the application"
	@echo "  make setup       - Run the setup script"
	@echo "  make test        - Run tests"
	@echo "  make clean       - Clean build artifacts"
	@echo "  make deb         - Build Debian package"
	@echo "  make snap        - Build Snap package"
	@echo "  make dev         - Build with race detector"
	@echo "  make build-all   - Build for multiple architectures"
	@echo "  make fmt         - Format code"
	@echo "  make lint        - Lint code"
	@echo "  make deps        - Update dependencies"
	@echo "  make help        - Show this help message"
	@echo ""
	@echo "Quick start:"
	@echo "  1. make build"
	@echo "  2. make setup"
	@echo "  3. make run"
	@echo ""
	@echo "Or install system-wide:"
	@echo "  1. make install"
	@echo "  2. cli-assistant"
