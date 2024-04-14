include .env
export

# Переменные
PROJECT_NAME := banner-service
DOCKER_IMAGE := $(PROJECT_NAME)
DOCKER_CONTAINER := $(PROJECT_NAME)-container

# Команды
.PHONY: all build run test clean docker-build docker-run docker-stop

all: build

MIGRATE_CMD=goose postgres "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_PORT)?sslmode=disable" up

migrate-up:
	cd ./migrations && $(MIGRATE_CMD)
build:
	go build -o bin/$(PROJECT_NAME) ./cmd/main.go

run:
	go run ./cmd/main.go

test:
	go test -v ./...

clean:
	go clean
	rm -rf bin

docker-build:
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	docker run -d --name $(DOCKER_CONTAINER) -p 8080:8080 $(DOCKER_IMAGE)

docker-stop:
	docker stop $(DOCKER_CONTAINER)
	docker rm $(DOCKER_CONTAINER)

# Цели по умолчанию
.DEFAULT_GOAL := build