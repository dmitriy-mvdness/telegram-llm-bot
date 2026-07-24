package service

import (
	"context"

	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/config"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/model"
)

type LLM interface {
	Chat(ctx context.Context, messages []model.Message, options config.LLMOptions) (string, error)
}
