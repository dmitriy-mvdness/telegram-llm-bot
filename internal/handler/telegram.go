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

			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text:   resp,
			})
			if err != nil {
				log.Printf("failed to send message to chat %d: %v", chatID, err)
			}
			return
		}

		_, err := b.SendChatAction(ctx, &bot.SendChatActionParams{
			ChatID: chatID,
			Action: models.ChatActionTyping,
		})
		if err != nil {
			log.Printf("failed to send chat action to chat %d: %v", chatID, err)
		}

		statusMsg, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "⏳ Бот думает...",
		})
		if err != nil {
			log.Printf("failed to send status message to chat %d: %v", chatID, err)
			resp := h.Handle(chatID, text)
			if _, err := b.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: resp}); err != nil {
				log.Printf("failed to fallback send to chat %d: %v", chatID, err)
			}
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
			if _, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text:   resp,
			}); err != nil {
				log.Printf("failed to fallback send after edit fail to chat %d: %v", chatID, err)
			}
		}
	})
}
