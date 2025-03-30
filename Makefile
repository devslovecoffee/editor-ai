.PHONY: build test clean lint run help

BINARY_NAME=editor-ai
GO=go
GOFLAGS=-ldflags="-s -w"

# Determine platform-specific binary name
ifeq ($(OS),Windows_NT)
	BINARY_NAME := $(BINARY_NAME).exe
	PLATFORM := windows-amd64
else
	UNAME_S := $(shell uname -s)
	UNAME_M := $(shell uname -m)
	ifeq ($(UNAME_S),Linux)
		PLATFORM := linux-amd64
	endif
	ifeq ($(UNAME_S),Darwin)
		ifeq ($(UNAME_M),arm64)
			PLATFORM := darwin-arm64
		else
			PLATFORM := darwin-amd64
		endif
	endif
endif

help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  build       Build the application"
	@echo "  test        Run tests"
	@echo "  clean       Remove build artifacts"
	@echo "  lint        Run linters"
	@echo "  run         Run the application (requires DIR and API_KEY)"
	@echo "  release     Build for all platforms"
	@echo "  help        Show this help"

build:
	$(GO) build $(GOFLAGS) -o $(BINARY_NAME) .

test:
	$(GO) test ./... -v

clean:
	$(GO) clean
	rm -f $(BINARY_NAME)
	rm -rf dist/

lint:
	$(GO) vet ./...
	@if command -v golint > /dev/null; then \
		golint ./...; \
	else \
		echo "golint is not installed"; \
	fi

run:
ifndef DIR
	$(error DIR is required. Usage: make run DIR="/path/to/files" API_KEY="your-api-key" [GLOB="pattern"])
endif
ifndef API_KEY
	$(error API_KEY is required. Usage: make run DIR="/path/to/files" API_KEY="your-api-key" [GLOB="pattern"])
endif
	$(GO) run main.go --dir "$(DIR)" --api_key "$(API_KEY)" $(if $(GLOB),--glob "$(GLOB)")

release: clean
	mkdir -p dist
	# Build for Linux
	GOOS=linux GOARCH=amd64 $(GO) build $(GOFLAGS) -o dist/$(BINARY_NAME)-linux-amd64 .
	# Build for macOS (Intel)
	GOOS=darwin GOARCH=amd64 $(GO) build $(GOFLAGS) -o dist/$(BINARY_NAME)-darwin-amd64 .
	# Build for macOS (M1/Apple Silicon)
	GOOS=darwin GOARCH=arm64 $(GO) build $(GOFLAGS) -o dist/$(BINARY_NAME)-darwin-arm64 .
	# Build for Windows
	GOOS=windows GOARCH=amd64 $(GO) build $(GOFLAGS) -o dist/$(BINARY_NAME)-windows-amd64.exe .
	# Create checksums
	cd dist && sha256sum * > checksums.txt

default: build 