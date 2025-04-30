# Makefile для itbookworm (Go + PostgreSQL)

.PHONY: build run test migrate-up migrate-down clean

# Переменные
APP_NAME := itbookworm
DB_URL := postgres://user:password@localhost:5432/itbookworm?sslmode=disable
MIGRATIONS_DIR := ./migrations

# Сборка приложения
build:
	go build -o bin/$(APP_NAME) ./cmd/api

# Запуск приложения
run: build
	./bin/$(APP_NAME)

# Запуск тестов
test:
	go test -v ./...

# Применить миграции БД
migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

# Откатить последнюю миграцию
migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down

# Очистка бинарников
clean:
	rm -rf bin/

# Запуск с Docker Compose
docker-up:
	docker-compose up --build -d

# Остановка Docker Compose
docker-down:
	docker-compose down

# Применить миграции в Docker
docker-migrate-up:
	docker-compose exec app migrate -path $(MIGRATIONS_DIR) -database "$$DB_URL" up

# Линтинг Go-кода
lint:
	golangci-lint run ./...

# Форматирование кода
fmt:
	go fmt ./...