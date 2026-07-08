package service

type LLM interface {
	Generate(prompt string) (string, error)
}
