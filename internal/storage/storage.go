package storage

import "github.com/dmitriy-mvdness/telegram-llm-bot/internal/model"

type MessageStore interface {
	Add(chatID int64, msg model.Message) error
	Get(chatID int64) ([]model.Message, error)
	Clear(chatID int64) error
}
