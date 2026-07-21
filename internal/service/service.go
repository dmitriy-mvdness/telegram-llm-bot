package service

import (
	"context"
	"fmt"
	"log"

	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/model"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/storage"
)

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

func (s *Service) Process(ctx context.Context, chatID int64, inputText string) string {
	err := s.store.Add(chatID, model.Message{
		Role:    "user",
		Content: inputText,
	})
	if err != nil {
		log.Printf("failed to save message: %v", err)

		return "❌ Не удалось получить ответ. Попробуйте позже"
	}

	history, err := s.store.Get(chatID)
	if err != nil {
		log.Printf("failed to get message history: %v", err)

		return "❌ Не удалось получить ответ. Попробуйте позже"
	}

	prompt, err := s.GetUserPrompt(chatID)
	if err != nil {
		log.Println(err)

		return "❌ Не удалось получить ответ. Попробуйте позже"
	}

	messages := append(
		[]model.Message{
			{
				Role:    "system",
				Content: prompt.Content,
			},
		},
		history...,
	)

	resp, err := s.llm.Chat(ctx, messages)
	if err != nil {
		log.Printf("failed to generate response: %v", err)

		return "❌ Не удалось получить ответ. Попробуйте позже"
	}

	if err := s.store.Add(chatID, model.Message{
		Role:    "assistant",
		Content: resp,
	}); err != nil {
		log.Printf("failed to save assistant message: %v", err)
	}

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

func (s *Service) EnsureUser(chatID int64) error {
	return s.user.Ensure(chatID)
}
