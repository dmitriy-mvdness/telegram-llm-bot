-include .env
export

.PHONY: run build clean tidy

ifeq ($(OS),Windows_NT)
    MKDIR = if not exist bin mkdir bin
    RMDIR = if exist bin rmdir /s /q bin
else
    MKDIR = mkdir -p bin
    RMDIR = rm -rf bin
endif

run:
	@echo Запуск программы...
	@go run ./cmd/app

build:
	@echo Сборка программы...
	@$(MKDIR)
	@go build -o bin/app ./cmd/app

clean:
	@echo Очистка...
	@$(RMDIR)

tidy:
	@echo Обновление зависимостей...
	@go mod tidy