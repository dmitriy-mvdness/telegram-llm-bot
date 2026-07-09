package service

import (
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
	llm    LLM
	memory *storage.Memory
}

func New(llm LLM) *Service {
	return &Service{
		llm:    llm,
		memory: storage.NewMemory(),
	}
}

func (s *Service) Process(userID, inputText string) string {
	s.memory.Add(userID, model.Message{
		Role:    "user",
		Content: inputText,
	})

	history := s.memory.Get(userID)

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

	s.memory.Add(userID, model.Message{
		Role:    "assistant",
		Content: resp,
	})

	return resp
}

func (s *Service) ClearMemory(userID string) {
	s.memory.Clear(userID)
}
