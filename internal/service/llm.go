package service

type LLM interface {
	Chat(messages []Message) (string, error)
}
