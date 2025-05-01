.PHONY: build test run clean

APP_NAME := itbookworm
DB_URL := postgres://user:password@localhost:5432/itbookworm?sslmode=disable
MIGRATIONS_DIR := ./migrations

build:
	go build -o bin/$(APP_NAME) ./cmd/api

run:
	DB_URL="postgres://user:password@localhost:5432/itbookworm?sslmode=disable" ./bin/itbookworm

clean:
	rm -rf bin/

docker-up:
	docker-compose up --build -d

docker-down:
	docker-compose down

docker-migrate-up:
	docker-compose exec app migrate -path $(MIGRATIONS_DIR) -database "$$DB_URL" up

test:
	go test -v -short ./internal/models/...