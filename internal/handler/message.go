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

		// Удаляем сообщение пользователя,
		// так как предыдущий запрос ещё выполняется
		if _, err := b.DeleteMessage(
			ctx,
			&bot.DeleteMessageParams{
				ChatID:    chatID,
				MessageID: messageID,
			},
		); err != nil {
			log.Printf("failed to delete busy user message: %v", err)
		}

		// Проверяем, отправляли ли уже уведомление
		if h.busy.GetNoticeMessage(chatID) == 0 {

			msg, err := b.SendMessage(
				ctx,
				&bot.SendMessageParams{
					ChatID: chatID,
					Text:   "❗ Подождите, я ещё обрабатываю предыдущий запрос",
				},
			)

			if err != nil {
				log.Printf("failed to send busy message: %v", err)
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

		resp := h.svc.Process(chatID, text)

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

	resp := h.svc.Process(chatID, text)

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
