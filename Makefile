BINARY_NAME = hexlet-path-size

.PHONY: install build test lint clean run

install:
	go mod tidy

build:
	go build -o bin/$(BINARY_NAME) ./cmd/$(BINARY_NAME)

test:
	go test -v ./...

lint:
	golangci-lint run

clean:
	rm -rf bin/ dist/

run: build
	./bin/$(BINARY_NAME) $(ARGS)
