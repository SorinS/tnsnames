# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOLINT=bin/golangci-lint run
BINARY_NAME=tnsnames

DATE=$(shell date +%Y%m%d_%H%M%S)
VERSION=0.1.0
COMMIT=$(shell git rev-parse HEAD)
LDFLAGS=-ldflags "-X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.buildDate=$(DATE)"

# Compilation targets
#all: macosx-arm64 linux-amd64 windows-amd64
all: macosx-arm64 macosx-amd64

macosx-arm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-macosx-arm64.bin

macosx-amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-macosx-amd64.bin

linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-linux-amd64.bin

windows-amd64:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-windows-amd64.exe

# Testing and linting targets
test:
	$(GOTEST) -v ./...

lint:
	$(GOLINT) ./...

clean:
	$(GOCLEAN)
	rm -f bin/*.bin

