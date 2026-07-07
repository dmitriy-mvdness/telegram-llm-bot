package service

import (
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/config"
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
	llm    *OllamaClient
	memory *Memory
}

func New(cfg config.Config) *Service {
	return &Service{
		llm:    NewOllamaClient(cfg.Ollama),
		memory: NewMemory(),
	}
}

func (s *Service) Process(userID, inputText string) string {
	s.memory.Add(userID, Message{
		Role:    "user",
		Content: inputText,
	})

	history := s.memory.Get(userID)

	prompt := buildPrompt(history)

	resp, err := s.llm.Generate(prompt)
	if err != nil {
		return "Ошибка генерации ответа: " + err.Error()
	}

	s.memory.Add(userID, Message{
		Role:    "assistant",
		Content: resp,
	})

	return resp
}

func buildPrompt(history []Message) string {
	prompt := systemPrompt + "\n" +
		"История сообщений:\n"

	for _, msg := range history {
		if msg.Role == "user" {
			prompt += "User: " + msg.Content + "\n"
		} else {
			prompt += "Assistant: " + msg.Content + "\n"
		}
	}

	prompt += "Твой ответ: "

	return prompt
}
