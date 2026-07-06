package service

import (
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/config"
)

const defaultModel = "openchat:7b" // Модель Ollama по умолчанию

const systemPrompt = `
Ты — AI-ассистент.

ПРАВИЛА:
- Не повторяй приветствие в каждом ответе
- Не начинай каждый ответ с "Привет"
- Не добавляй лишние обращения по имени без причины
- Отвечай кратко и по делу
- Если информации нет — скажи "я не знаю"
- Не выдумывай факты
- Не давай медицинских, юридических или возрастных оценок как истину

СТИЛЬ:
- естественный разговор
- без повторов
`

type Service struct {
	llm    *OllamaClient
	memory *Memory
}

func New(cfg config.Config) *Service {
	return &Service{
		llm:    NewOllamaClient(cfg.OllamaHost, defaultModel),
		memory: NewMemory(),
	}
}

func (s *Service) Process(userID, inputText string) string {
	s.memory.Add(userID, Message{
		Role:    "user",
		Content: inputText,
	})

	history := s.memory.Get(userID)

	prompt := buildPrompt(history, inputText)

	resp, err := s.llm.Generate(prompt)
	if err != nil {
		return err.Error()
	}

	s.memory.Add(userID, Message{
		Role:    "user",
		Content: resp,
	})

	return resp
}

func buildPrompt(history []Message, userText string) string {
	promt := systemPrompt + "\n\n"

	for _, msg := range history {
		if msg.Role == "user" {
			promt += "User: " + msg.Content + "\n"
		} else {
			promt += "Assistant: " + msg.Content + "\n"
		}
	}

	promt += "User: " + userText + "\nAssistant:"
	return promt
}
