package service

import (
	"context"
	"fmt"

	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/config"
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

func (s *Service) Process(ctx context.Context, chatID int64, inputText string) (string, error) {
	err := s.store.Add(chatID, model.Message{
		Role:    "user",
		Content: inputText,
	})
	if err != nil {
		return "", fmt.Errorf("failed to save message: %w", err)
	}

	history, err := s.store.Get(chatID)
	if err != nil {
		return "", fmt.Errorf("failed to get message history: %w", err)
	}

	prompt, err := s.GetUserPrompt(chatID)
	if err != nil {
		return "", fmt.Errorf("get user prompt: %w", err)
	}

	promptCfg, ok := config.PromptConfigs[prompt.Name]

	if !ok {
		promptCfg = config.PromptConfigs["Обычный"]
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

	resp, err := s.llm.Chat(ctx, messages, promptCfg.Options)
	if err != nil {
		return "", fmt.Errorf("failed to generate llm response: %w", err)
	}

	if err := s.store.Add(chatID, model.Message{
		Role:    "assistant",
		Content: resp,
	}); err != nil {
		return "", fmt.Errorf("failed to save assistant message: %w", err)
	}

	return resp, nil
}

func (s *Service) ClearHistory(chatID int64) error {
	if err := s.store.Clear(chatID); err != nil {
		return fmt.Errorf("clear history for chat %d: %w", chatID, err)
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
	if err := s.user.Ensure(chatID); err != nil {
		return fmt.Errorf("ensure user %d: %w", chatID, err)
	}

	return nil
}
