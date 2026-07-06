package service

import (
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/config"
)

type Service struct {
	llm *OllamaClient
}

const defaultModel = "openchat:7b" // Модель Ollama по умолчанию

func New(cfg config.Config) *Service {
	return &Service{
		llm: NewOllamaClient(cfg.OllamaHost, defaultModel),
	}
}

func (s *Service) Process(inputText string) string {
	resp, err := s.llm.Generate(inputText)
	if err != nil {
		return err.Error()
	}

	return resp
}
