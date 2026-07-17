package storage

import "github.com/dmitriy-mvdness/telegram-llm-bot/internal/model"

type UserStore interface {
	Ensure(chatID int64) error
	GetUserPrompt(chatID int64) (model.Prompt, error)
	UpdatePrompt(chatID int64, promptID int) error
}
