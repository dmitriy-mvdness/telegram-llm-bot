package handler

import (
	"context"
	"log"
	"strconv"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (h *Handler) Register(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypeContains, func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message == nil {
			return
		}

		statusMsg, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Бот печатает...",
		})
		if err != nil {
			log.Println(err)
		}

		userID := strconv.FormatInt(update.Message.Chat.ID, 10)

		resp := h.Handle(userID, update.Message.Text)

		_, err = b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:    update.Message.Chat.ID,
			MessageID: statusMsg.ID,
			Text:      resp,
		},
		)
		if err != nil {
			log.Println(err)
		}
	})
}
