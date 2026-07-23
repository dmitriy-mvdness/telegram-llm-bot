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
	messageID := update.Message.ID

	if !h.busy.TryLock(chatID) {

		if _, err := b.DeleteMessage(ctx, &bot.DeleteMessageParams{
			ChatID:    chatID,
			MessageID: messageID,
		}); err != nil {
			log.Printf("failed to delete busy user message: %v", err)
		}

		if h.busy.TryReserveNotice(chatID) {
			msg, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text:   "❗ Подождите, я ещё обрабатываю предыдущий запрос",
			})

			if err != nil {
				log.Printf("failed to send busy message: %v", err)
				h.busy.ReleaseNoticeReservation(chatID)
				return
			}

			h.busy.SetNoticeMessage(chatID, msg.ID)
		}

		return
	}

	defer func() {
		noticeID := h.busy.GetNoticeMessage(chatID)

		h.busy.Unlock(chatID)

		if noticeID != 0 {
			if _, err := b.DeleteMessage(
				ctx,
				&bot.DeleteMessageParams{
					ChatID:    chatID,
					MessageID: noticeID,
				},
			); err != nil {
				log.Printf("failed to delete notice message: %v", err)
			}
		}
	}()

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
		log.Printf("failed to send typing action: %v", err)
	}

	statusMsg, err := b.SendMessage(
		ctx,
		&bot.SendMessageParams{
			ChatID: chatID,
			Text:   "⏳ Бот думает...",
		},
	)

	if err != nil {
		log.Printf("failed to send status message: %v", err)

		resp, err := h.svc.Process(ctx, chatID, text)

		if err != nil {
			log.Printf(
				"process failed chat=%d: %v",
				chatID,
				err,
			)

			resp = defaultErrorMessage
		}

		if _, err := b.SendMessage(
			ctx,
			&bot.SendMessageParams{
				ChatID: chatID,
				Text:   resp,
			},
		); err != nil {
			log.Printf("failed fallback send: %v", err)
		}

		return
	}

	h.processingMessages.Store(chatID, statusMsg.ID)
	defer h.processingMessages.Delete(chatID)

	if h.isShuttingDown.Load() {
		b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:    chatID,
			MessageID: statusMsg.ID,
			Text:      "⚠️ Сервис временно недоступен",
		})
		return
	}

	resp, err := h.svc.Process(ctx, chatID, text)
	if err != nil {
		log.Printf("process failed chat=%d: %v", chatID, err)

		resp = defaultErrorMessage
	}

	_, err = b.EditMessageText(
		ctx,
		&bot.EditMessageTextParams{
			ChatID:    chatID,
			MessageID: statusMsg.ID,
			Text:      resp,
		},
	)

	if err != nil {
		log.Printf("failed to edit message: %v", err)

		if _, err := b.SendMessage(
			ctx,
			&bot.SendMessageParams{
				ChatID: chatID,
				Text:   resp,
			},
		); err != nil {
			log.Printf("failed fallback send after edit: %v", err)
		}
	}
}

func (h *Handler) Shutdown(ctx context.Context, b *bot.Bot) {
	h.isShuttingDown.Store(true)

	h.processingMessages.Range(func(key, value interface{}) bool {
		chatID := key.(int64)
		msgID := value.(int)

		_, err := b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:    chatID,
			MessageID: msgID,
			Text:      "⚠️ Произошла ошибка, попробуйте позже",
		})

		if err != nil {
			log.Printf("failed to edit processing message for chat %d: %v", chatID, err)
		}

		return true
	})
}
