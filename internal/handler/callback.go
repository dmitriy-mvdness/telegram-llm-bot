package handler

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *Handler) handleCallback(
	ctx context.Context,
	b *bot.Bot,
	update *models.Update,
) {
	if update.CallbackQuery == nil {
		return
	}

	_, err := b.AnswerCallbackQuery(
		ctx,
		&bot.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
		},
	)
	if err != nil {
		log.Printf("callback answer error: %v", err)
	}

	data := update.CallbackQuery.Data
	chatID := update.CallbackQuery.Message.Message.Chat.ID

	switch data {
	case "clear_history":
		if err := h.svc.ClearHistory(chatID); err != nil {
			log.Printf("failed to clear history: %v", err)

			if _, err := b.SendMessage(
				ctx,
				&bot.SendMessageParams{
					ChatID: chatID,
					Text:   "❌ Не удалось очистить историю",
				},
			); err != nil {
				log.Printf("failed to send clear history error message: %v", err)
			}
			return
		}
		if _, err := b.SendMessage(
			ctx,
			&bot.SendMessageParams{
				ChatID: chatID,
				Text:   "✅ История очищена",
			},
		); err != nil {
			log.Printf("failed to send history cleared message: %v", err)
		}
	case "system_prompts":
		if _, err := b.SendMessage(
			ctx,
			&bot.SendMessageParams{
				ChatID: chatID,
				Text: ` 🎭 Роль ассистента
				
Какой стиль ответов вам подходит?

Выберите режим:
				`,
				ReplyMarkup: AssistantRoleKeyboard(),
			},
		); err != nil {
			log.Printf("failed to send prompts message: %v", err)
		}

	default:
		log.Printf("unknown callback: %s", data)
	}
}
