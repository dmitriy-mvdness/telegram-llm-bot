package service

import (
	"fmt"

	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/model"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/storage"
)

const systemPrompt = `
Ты — AI-ассистент.
Пиши только на русском языке.
Отвечай естественно и по делу.
Если не знаешь — скажи об этом.
Не выдумывай информацию.
Пиши естественно, без лишних приветствий и повторов.
`

type Service struct {
	llm   LLM
	store storage.MessageStore
	user  storage.UserStore
}

func New(llm LLM, store storage.MessageStore, user storage.UserStore) *Service {
	return &Service{
		llm:   llm,
		store: store,
		user:  user,
	}
}

func (s *Service) Process(chatID int64, inputText string) string {
	if err := s.user.Ensure(chatID); err != nil {
		return "Ensure user failed: " + err.Error()
	}

	err := s.store.Add(chatID, model.Message{
		Role:    "user",
		Content: inputText,
	})
	if err != nil {
		return "Ошибка сохранения сообщения: " + err.Error()
	}

	history, err := s.store.Get(chatID)
	if err != nil {
		return "Ошибка получения истори сообщений: " + err.Error()
	}

	messages := append(
		[]model.Message{
			{
				Role:    "system",
				Content: systemPrompt,
			},
		},
		history...,
	)

	resp, err := s.llm.Chat(messages)
	if err != nil {
		return "Ошибка генерации ответа: " + err.Error()
	}

	s.store.Add(chatID, model.Message{
		Role:    "assistant",
		Content: resp,
	})

	return resp
}

func (s *Service) ClearHistory(chatID int64) error {
	if err := s.store.Clear(chatID); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetUserPrompt(chatID int64) (model.Prompt, error) {
	prompt, err := s.user.GetUserPrompt(chatID)
	if err != nil {
		return model.Prompt{}, fmt.Errorf("failed to get current prompt for chat %d: %w", chatID, err)
	}

	return prompt, nil
}

func (s *Service) UpdateUserPrompt(chatID int64, promptID int) error {
	if err := s.user.UpdatePrompt(chatID, promptID); err != nil {
		return fmt.Errorf("failed to update prompt for chat %d: %w", chatID, err)
	}

	return nil
}
