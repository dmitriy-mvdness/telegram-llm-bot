package storage

import "github.com/dmitriy-mvdness/telegram-llm-bot/internal/model"

type MessageStore interface {
	Add(userID string, msg model.Message)
	Get(userID string) []model.Message
	Clear(userID string)
}