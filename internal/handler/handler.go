package handler

import (
	"strings"

	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/service"
)

type Handler struct {
	svc *service.Service
}

// Создать новый обработчик
func New(svc *service.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

// Обработка сообщений
func (h *Handler) Handle(userID, inputText string) string {
	if inputText == "" {
		return ""
	}

	if h.isCommand(inputText) {
		return h.handleCommand(userID, inputText)
	}

	return h.handleChat(userID, inputText)
}

// Проверка на тип сообщения
func (h *Handler) isCommand(inputText string) bool {
	return len(inputText) > 0 && inputText[0] == '/'
}

// Если команда
func (h *Handler) handleCommand(userID, inputText string) string {
	inputText = strings.TrimSpace(inputText)

	switch inputText {
	case "/start":
		return "👋 Привет!\n" +
			"Я AI-Ассистент. Задай любой вопрос, и я постараюсь помочь :-]"

	case "/help":
		return "Доступные команды:\n" +
			"/start - приветствие\n" +
			"/help - помощь\n" +
			"/clear - очистить историю диалога\n\n" +
			"Или просто напиши свой вопрос!"

	case "/clear":
		h.svc.ClearMemory(userID)
		return "История диалога очищена!"

	default:
		return "Неизвестная команда!"
	}
}

// Если обычное сообщение
func (h *Handler) handleChat(userID string, inputText string) string {
	return h.svc.Process(userID, inputText)
}
