package handler

import (
	"context"
	"fmt"
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

	if !h.ensureUser(chatID) {
		return
	}

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
		prompt, err := h.svc.GetUserPrompt(chatID)
		if err != nil {
			log.Println(err)
		}

		if _, err := b.SendMessage(
			ctx,
			&bot.SendMessageParams{
				ChatID: chatID,
				Text: fmt.Sprintf(`🎭 Роль ассистента

Текущая роль: %s

Режимы:
🤖 Обычный — универсальный стиль
⚡ Краткий — только главное
🎓 Экспертный — точные и строгие ответы
🧠 Наводящий — через вопросы и рассуждения
📚 Подробный — объяснения с примерами

Выберите стиль ответов:`, prompt.DisplayName),
				ReplyMarkup: AssistantRoleKeyboard(),
			},
		); err != nil {
			log.Printf("failed to send prompts message: %v", err)
		}
	case "prompt_default":
		if err := h.svc.UpdateUserPrompt(chatID, 1); err != nil {
			log.Println(err)
		}

		if _, err := b.SendMessage(
			ctx,
			&bot.SendMessageParams{
				ChatID: chatID,
				Text:   PromptSelectedMessage("🤖 Обычный"),
			},
		); err != nil {
			log.Printf("failed to send prompt activate message: %v", err)
		}
	case "prompt_concise":
		if err := h.svc.UpdateUserPrompt(chatID, 2); err != nil {
			log.Println(err)
		}

		if _, err := b.SendMessage(
			ctx,
			&bot.SendMessageParams{
				ChatID: chatID,
				Text:   PromptSelectedMessage("⚡ Краткий"),
			},
		); err != nil {
			log.Printf("failed to send prompt activate message: %v", err)
		}
	case "prompt_academic":
		if err := h.svc.UpdateUserPrompt(chatID, 3); err != nil {
			log.Println(err)
		}

		if _, err := b.SendMessage(
			ctx,
			&bot.SendMessageParams{
				ChatID: chatID,
				Text:   PromptSelectedMessage("🎓 Экспертный"),
			},
		); err != nil {
			log.Printf("failed to send prompt activate message: %v", err)
		}
	case "prompt_provocative":
		if err := h.svc.UpdateUserPrompt(chatID, 4); err != nil {
			log.Println(err)
		}

		if _, err := b.SendMessage(
			ctx,
			&bot.SendMessageParams{
				ChatID: chatID,
				Text:   PromptSelectedMessage("🧠 Наводящий"),
			},
		); err != nil {
			log.Printf("failed to send prompt activate message: %v", err)
		}
	case "prompt_encyclopedic":
		if err := h.svc.UpdateUserPrompt(chatID, 5); err != nil {
			log.Println(err)
		}

		if _, err := b.SendMessage(
			ctx,
			&bot.SendMessageParams{
				ChatID: chatID,
				Text:   PromptSelectedMessage("📚 Подробный"),
			},
		); err != nil {
			log.Printf("failed to send prompt activate message: %v", err)
		}
	case "stop_generation":
		h.generation.Cancel(chatID)

		if msgID, ok := h.processingMessages.Load(chatID); ok {

			_, err := b.DeleteMessage(
				ctx,
				&bot.DeleteMessageParams{
					ChatID:    chatID,
					MessageID: msgID.(int),
				},
			)

			if err != nil {
				log.Printf("failed to delete processing message: %v", err)
			}
		}

	default:
		log.Printf("unknown callback: %s", data)
	}
}
