package config

type Config struct {
	TelegramClient string
	DatabaseURL    string
	OllamaHost     string
}

func Load() Config {
	return Config{}
}
