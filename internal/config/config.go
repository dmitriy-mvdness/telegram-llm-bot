package config

import "os"

type Config struct {
	TelegramClient string
	DatabaseURL    string
	OllamaHost     string
}

func Load() Config {
	return Config{
		TelegramClient: os.Getenv("TELEGRAM_TOKEN"),
		OllamaHost:     os.Getenv("OLLAMA_BASE_URL"),
	}
}
