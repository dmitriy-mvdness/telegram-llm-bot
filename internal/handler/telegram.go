package handler

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *Handler) Register(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypeContains, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message == nil {
			return
		}

		chatID := update.Message.Chat.ID
		text := update.Message.Text

		if h.isCommand(text) {
			resp := h.Handle(chatID, text)

			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text:   resp,
			})
			return
		}

		_, err := b.SendChatAction(ctx, &bot.SendChatActionParams{
			ChatID: chatID,
			Action: models.ChatActionTyping,
		})
		if err != nil {
			log.Printf("failed to send chat action: %v", err)
		}

		statusMsg, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "⏳ Бот думает...",
		})
		if err != nil {
			log.Printf("failed to send status message: %v", err)
			resp := h.Handle(chatID, text)
			b.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: resp})
			return
		}

		resp := h.Handle(chatID, text)

		_, err = b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:    chatID,
			MessageID: statusMsg.ID,
			Text:      resp,
		})
		if err != nil {
			log.Printf("failed to edit message: %v", err)
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text:   resp,
			})
		}
	})
}
