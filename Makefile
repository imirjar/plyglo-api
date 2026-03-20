# Define project variables
PROJECT_NAME := poliglotim-api
GO_FILES := $(shell find . -name '*.go' -print)

.PHONY: docs
docs:
	swag init -g internal/gateway/http/server.go -o docs/

# Default target: build the application
.PHONY: all
all: build

# Build target: compiles the Go application
.PHONY: build
build:
	go build -o ./bin/$(PROJECT_NAME) ./cmd/main.go
# 	CGO_ENABLED=0 GOOS=linux go build -o bin/$(BIN_NAME) $(APP_FILE)
# 	GOOS=linux GOARCH=amd64 go build -o $(BIN_NAME) $(APP_FILE)

# Test target: runs all unit tests
.PHONY: test
test:
	go test -v ./...

# Clean target: removes compiled binaries and other artifacts
.PHONY: clean
clean:
	rm -f ./bin/$(PROJECT_NAME)
	go clean

# Format target: formats Go source code using gofmt
.PHONY: fmt
fmt:
	gofmt -w $(GO_FILES)

# Lint target: runs golint for code style checks
.PHONY: lint
lint:
	golint ./...

# Vet target: runs go vet for static analysis
.PHONY: vet
vet:
	go vet ./...

# Install dependencies target
.PHONY: deps
deps:
	go mod tidy
	go mod download

# Help target: displays available commands
.PHONY: help
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  all     - Builds the application (default)"
	@echo "  build   - Compiles the Go application"
	@echo "  test    - Runs all unit tests"
	@echo "  clean   - Removes compiled binaries and other artifacts"
	@echo "  fmt     - Formats Go source code"
	@echo "  lint    - Runs golint for code style checks"
	@echo "  vet     - Runs go vet for static analysis"
	@echo "  deps    - Installs Go module dependencies"
	@echo "  help    - Displays this help message"
