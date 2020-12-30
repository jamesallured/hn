GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=hn

all: clean build
build:
	$(GOBUILD) -o bin/$(BINARY_NAME) -v ./cmd/$(BINARY_NAME)
clean:
	$(GOCLEAN)
	rm -rf bin