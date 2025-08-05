ci:
	docker compose -f docker-compose.yml down -v
	docker compose -f docker-compose.yml build
	docker compose -f docker-compose.yml run --rm app make setup
	docker compose -f docker-compose.yml up --abort-on-container-exit

compose-setup: compose-build compose-app-setup

compose-build:
	docker compose build

compose-bash:
	docker compose run --rm app bash

compose-test:
	docker compose -f docker-compose.yml up --abort-on-container-exit

compose-lint:
	docker compose run --rm app make lint

compose:
	docker compose up

compose-down:
	docker compose down -v --remove-orphans

compose-app-setup:
	docker compose run --rm app make setup

setup:
	cd code && go mod tidy
	cd tests && go mod tidy

update-deps:
	@echo "📦 Updating dependencies in code/"
	cd code && go get -u ./... && go mod tidy
	@echo "📦 Updating dependencies in tests/"
	cd tests && go get -u ./... && go mod tidy

test:
	go test -v ./tests

lint:
	golangci-lint run tests/... code/...

code-test:
	go test -v ./...
