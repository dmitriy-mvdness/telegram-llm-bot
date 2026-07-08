package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/config"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/handler"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/service"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
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

	llm := service.NewOllamaClient(cfg.Ollama)

	svc := service.New(llm)

	h := handler.New(svc)

	b, err := bot.New(token)
	if err != nil {
		log.Fatal(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypeContains, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message == nil {
			return
		}

		userID := strconv.FormatInt(update.Message.Chat.ID, 10)

		resp := h.Handle(userID, update.Message.Text)

		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   resp,
		})
		if err != nil {
			log.Println(err)
		}
	})

	log.Println("Bot started!")
	b.Start(ctx)
}
