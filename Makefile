-include .env
export

.PHONY: run build clean

run:
	@echo Запуск программы...
	@go mod tidy
	@go run ./cmd/app

build:
	@echo Сборка программы...
	@go mod tidy
	@mkdir bin
	@go build -o bin/app ./cmd/app

clean:
	@go clean
	@if exist bin rmdir /s /q bin