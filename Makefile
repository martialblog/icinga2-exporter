.PHONY: build

VERSION := $(shell git rev-parse HEAD)

build:
	go build -ldflags "-X main.build=$(VERSION)"
fmt:
	go fmt *.go
test:
	go test -v
