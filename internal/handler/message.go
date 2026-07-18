package handler

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *Handler) handleMessage(
	ctx context.Context,
	b *bot.Bot,
	update *models.Update,
) {
	chatID := update.Message.Chat.ID
	text := update.Message.Text

	if !h.ensureUser(chatID) {
		return
	}

	if _, err := b.SendChatAction(
		ctx,
		&bot.SendChatActionParams{
			ChatID: chatID,
			Action: models.ChatActionTyping,
		},
	); err != nil {
		log.Printf("failed to type action: %v", err)
	}

	statusMsg, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   "⏳ Бот думает...",
	})
	if err != nil {
		log.Printf("failed to send status message to chat %d: %v", chatID, err)
		resp := h.svc.Process(chatID, text)
		if _, err := b.SendMessage(ctx, &bot.SendMessageParams{ChatID: chatID, Text: resp}); err != nil {
			log.Printf("failed to fallback send to chat %d: %v", chatID, err)
		}
		return
	}

	resp := h.svc.Process(chatID, text)

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
}
