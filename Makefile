.PHONY: build

build:
	go build
fmt:
	go fmt *.go
test:
	go test -v
