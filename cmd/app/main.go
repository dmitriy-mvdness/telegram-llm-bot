package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/busy"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/config"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/database"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/handler"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/llm"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/service"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/storage/sqlite"
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

	db, err := database.Open(cfg.DatabaseDriver, cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database opened successfully")

	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// Костыль, тк БД после миграций закрывается
	db, err = database.Open(cfg.DatabaseDriver, cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

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

	messageStore := sqlite.New(db)
	userStore := sqlite.NewUserStore(db)

	busyManager := busy.NewManager()

	svc := service.New(llm, messageStore, userStore)

	h := handler.New(svc, busyManager)

	b, err := bot.New(token)
	if err != nil {
		log.Fatal(err)
	}

	h.Register(b)

	log.Println("Bot started!")
	b.Start(ctx)

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	h.Shutdown(shutdownCtx, b)

	db.Close()

	log.Println("Shutdown complete")
}
