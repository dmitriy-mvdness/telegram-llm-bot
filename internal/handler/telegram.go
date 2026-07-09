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

		userID := strconv.FormatInt(update.Message.Chat.ID, 10)

		if h.isCommand(update.Message.Text) {
			resp := h.Handle(
				userID,
				update.Message.Text,
			)

			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: userID,
				Text:   resp,
			})
			return
		}

		statusMsg, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Бот печатает...",
		})
		if err != nil {
			log.Println(err)
		}

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
