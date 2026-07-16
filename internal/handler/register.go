package handler

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *Handler) Register(b *bot.Bot) {
	b.RegisterHandler(
		bot.HandlerTypeMessageText,
		"",
		bot.MatchTypeContains,
		h.handleUpdate,
	)

	b.RegisterHandler(
		bot.HandlerTypeCallbackQueryData,
		"",
		bot.MatchTypePrefix,
		h.handleCallback,
	)
}

func (h *Handler) handleUpdate(
	ctx context.Context,
	b *bot.Bot,
	update *models.Update,
) {
	if update.Message == nil {
		return
	}

	text := update.Message.Text

	if text == "" {
		return
	}

	if h.isCommand(text) {
		h.handleCommand(ctx, b, update)
		return
	}

	h.handleMessage(ctx, b, update)
}
