package service

import "github.com/dmitriy-mvdness/telegram-llm-bot/internal/model"

type LLM interface {
	Chat(messages []model.Message) (string, error)
}
