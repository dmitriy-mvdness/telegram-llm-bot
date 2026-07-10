package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/config"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/db"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/handler"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/llm"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/service"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/storage"
	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_TOKEN is empty!")
	}

	cfg := config.Load()

	database, err := db.Open(cfg.DatabaseDriver, cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	llm := llm.NewOllamaClient(cfg.Ollama)

	log.Println("Checking Ollama...")
	if err := llm.Health(ctx); err != nil {
		log.Fatalf("Ollama health check failed: %v", err)
	}
	log.Println("Ollama is available.")

	log.Println("Checking LLM readiness...")
	if err := llm.Ready(ctx); err != nil {
		log.Fatalf("LLM health check failed: %v", err)
	}

	log.Println("LLM is ready.")

	memory := storage.NewMemory()

	svc := service.New(llm, memory)

	h := handler.New(svc)

	b, err := bot.New(token)
	if err != nil {
		log.Fatal(err)
	}

	h.Register(b)

	log.Println("Bot started!")
	b.Start(ctx)
}
