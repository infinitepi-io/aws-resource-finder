VERSION ?= $(shell git describe --tags --always)
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
BINARY_NAME=aws-resource-finder

build:
	@mkdir -p bin
	go build -v -o bin/$(BINARY_NAME)-$(GOOS)-$(GOARCH) -ldflags "-X main.Version=$(VERSION)" .
	@echo "Build complete: bin/$(BINARY_NAME)-$(GOOS)-$(GOARCH)"

build-all:
	GOOS=linux GOARCH=amd64 make build
	GOOS=linux GOARCH=arm64 make build
	GOOS=darwin GOARCH=amd64 make build
	GOOS=darwin GOARCH=arm64 make build
	GOOS=windows GOARCH=amd64 make build
	@echo "Build complete for all platforms"

clean:
	rm -rf bin/

.PHONY: build build-all clean