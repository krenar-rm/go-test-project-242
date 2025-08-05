.PHONY: install test

install:
	go mod tidy

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

update-deps:
	go get -u ./...
	go mod tidy

test:
	go test -v ./...
