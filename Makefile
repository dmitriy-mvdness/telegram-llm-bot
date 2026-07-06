include .env
export

run:
	@echo "Запуск программы..."
	go run .cmd/app/main.go

build:
	@echo "Сборка программы..."
	go build -o bin/app cmd/app/main.go