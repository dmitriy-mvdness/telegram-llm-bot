package handler

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *Handler) handleCommand(
	ctx context.Context,
	b *bot.Bot,
	update *models.Update,
) {
	chatID := update.Message.Chat.ID
	command := update.Message.Text

	if !h.ensureUser(chatID) {
		return
	}

	switch command {
	case "/start":
		_, err := b.SendMessage(
			ctx,
			&bot.SendMessageParams{
				ChatID: chatID,
				Text: `
👋 Привет!

Я AI-ассистент
Задай любой вопрос — постараюсь помочь

Больше возможностей и настройки: /help
				`,
			},
		)
		if err != nil {
			log.Printf("failed to send /start message %v", err)
		}
	case "/help":
		_, err := b.SendMessage(
			ctx,
			&bot.SendMessageParams{
				ChatID: chatID,
				Text: `
⌨️ Доступные команды:

/start — приветственное сообщение
/help — показать список доступных команд
/settings — настроить ассистента
				`,
			},
		)
		if err != nil {
			log.Printf("failed to send /help message %v", err)
		}
	case "/settings":
		_, err := b.SendMessage(
			ctx,
			&bot.SendMessageParams{
				ChatID: chatID,
				Text: `
⚙️ Настройки

Здесь можно управлять работой ассистента
					
Выберите действие:
				`,
				ReplyMarkup: SettingsKeyboard(),
			},
		)
		if err != nil {
			log.Printf("failed to send /settings message %v", err)
		}
	default:
		_, err := b.SendMessage(
			ctx,
			&bot.SendMessageParams{
				ChatID: chatID,
				Text:   "❓ Неизвестная команда",
			},
		)
		if err != nil {
			log.Printf("failed to send default message %v", err)
		}
	}

}
