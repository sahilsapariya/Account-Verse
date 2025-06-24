# Project variables
BINARY_NAME=server
MAIN_PATH=./main.go
BUILD_DIR=./build
COVERAGE_DIR=./coverage

# Go related variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=gofmt
GOLINT=golangci-lint

# Make targets phony
.PHONY: help build clean test test-coverage test-race test-verbose lint lint-fix format deps deps-update deps-verify gqlgen-generate run dev air docker-build docker-run install-tools setup check pre-commit all

# Default target
all: clean deps format lint test build

# Help target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Build targets
build: ## Build the application
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Build completed: $(BUILD_DIR)/$(BINARY_NAME)"

build-linux: ## Build for Linux
	@echo "Building for Linux..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-linux $(MAIN_PATH)

build-windows: ## Build for Windows
	@echo "Building for Windows..."
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-windows.exe $(MAIN_PATH)

build-mac: ## Build for macOS
	@echo "Building for macOS..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-mac $(MAIN_PATH)

build-all: build-linux build-windows build-mac ## Build for all platforms

# Clean targets
clean: ## Clean build artifacts
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)
	@rm -rf $(COVERAGE_DIR)
	@rm -f coverage.out coverage.html
	@echo "Clean completed"

# Test targets
test: ## Run tests
	@echo "Running tests..."
	$(GOTEST) ./test/... -v

test-all: ## Run all tests including packages without test files
	@echo "Running all tests..."
	$(GOTEST) ./... -v

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@mkdir -p $(COVERAGE_DIR)
	$(GOTEST) ./test/... -coverprofile=$(COVERAGE_DIR)/coverage.out
	$(GOCMD) tool cover -html=$(COVERAGE_DIR)/coverage.out -o $(COVERAGE_DIR)/coverage.html
	@echo "Coverage report generated: $(COVERAGE_DIR)/coverage.html"

test-race: ## Run tests with race detector
	@echo "Running tests with race detector..."
	$(GOTEST) ./test/... -race -v

test-verbose: ## Run tests with verbose output
	@echo "Running tests with verbose output..."
	$(GOTEST) ./test/... -v -count=1

test-bench: ## Run benchmark tests
	@echo "Running benchmark tests..."
	$(GOTEST) ./test/... -bench=. -benchmem

# Linting and formatting targets
lint: ## Run linter
	@echo "Running linter..."
	@export PATH=$$PATH:$$(go env GOPATH)/bin && $(GOLINT) run

lint-fix: ## Run linter with auto-fix
	@echo "Running linter with auto-fix..."
	@export PATH=$$PATH:$$(go env GOPATH)/bin && $(GOLINT) run --fix

format: ## Format code
	@echo "Formatting code..."
	$(GOFMT) -s -w .
	@export PATH=$$PATH:$$(go env GOPATH)/bin && goimports -w .

format-check: ## Check if code is formatted
	@echo "Checking code format..."
	@if [ -n "$$($(GOFMT) -s -l .)" ]; then \
		echo "Code is not formatted. Please run 'make format'"; \
		$(GOFMT) -s -l .; \
		exit 1; \
	fi

# Dependency management
deps: ## Download dependencies
	@echo "Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

deps-update: ## Update dependencies
	@echo "Updating dependencies..."
	$(GOGET) -u ./...
	$(GOMOD) tidy

deps-verify: ## Verify dependencies
	@echo "Verifying dependencies..."
	$(GOMOD) verify

deps-graph: ## Generate dependency graph
	@echo "Generating dependency graph..."
	$(GOCMD) mod graph

# GraphQL generation
gqlgen-generate: ## Generate GraphQL code
	@echo "Generating GraphQL code..."
	$(GOGET) github.com/99designs/gqlgen/codegen/config@v0.17.75
	$(GOGET) github.com/99designs/gqlgen/internal/imports@v0.17.75
	$(GOGET) github.com/99designs/gqlgen/api@v0.17.75
	$(GOGET) github.com/99designs/gqlgen@v0.17.75
	$(GOCMD) run github.com/99designs/gqlgen generate
	@echo "GraphQL code generation completed"

# Development targets
run: ## Run the application
	@echo "Running $(BINARY_NAME)..."
	$(GOCMD) run $(MAIN_PATH)

dev: ## Run in development mode with Air
	@echo "Starting development server with Air..."
	@export PATH=$$PATH:$$(go env GOPATH)/bin && air -c .air.toml

air: dev ## Alias for dev

# Docker targets
docker-build: ## Build Docker image
	@echo "Building Docker image..."
	docker build -t $(BINARY_NAME) .

docker-run: ## Run Docker container
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(BINARY_NAME)

# Installation and setup
install-tools: ## Install development tools
	@echo "Installing development tools..."
	@$(GOCMD) install github.com/99designs/gqlgen@v0.17.75
	@$(GOCMD) install github.com/air-verse/air@latest
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.55.2
	@$(GOCMD) install golang.org/x/tools/cmd/goimports@latest
	@echo "Development tools installed"

setup: install-tools deps ## Setup development environment
	@echo "Setting up development environment..."
	@echo "Development environment setup completed"

# Quality checks
check: format-check lint test ## Run all quality checks

pre-commit: format lint test ## Run pre-commit checks
	@echo "Pre-commit checks completed successfully"
	@echo "Staging changes..."
	cd .. && git add .

# Database targets
db-migrate: ## Run database migrations
	@echo "Running database migrations..."
	# Add your migration command here

db-seed: ## Seed database with test data
	@echo "Seeding database..."
	# Add your seeding command here

# Security targets
security-scan: ## Run security scan
	@echo "Running security scan..."
	@$(GOCMD) install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest
	@export PATH=$$PATH:$$(go env GOPATH)/bin && gosec ./...

# Performance targets
profile-cpu: ## Profile CPU usage
	@echo "Profiling CPU usage..."
	$(GOTEST) ./test/... -cpuprofile=cpu.prof
	$(GOCMD) tool pprof cpu.prof

profile-mem: ## Profile memory usage
	@echo "Profiling memory usage..."
	$(GOTEST) ./test/... -memprofile=mem.prof
	$(GOCMD) tool pprof mem.prof

# Release targets
release: clean deps format lint test build-all ## Prepare release
	@echo "Release preparation completed"

# Monitoring targets
health-check: ## Check application health
	@echo "Checking application health..."
	@curl -f http://localhost:8080/health || echo "Health check failed"
