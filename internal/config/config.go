package config

const (
	DefaultOllamaURL   = "http://localhost:11434"
	DefaultOllamaModel = "deepseek-r1:7b"

	DefaultNumCtx        = 2048
	DefaultNumPredict    = 300
	DefaultTemperature   = 0.3
	DefaultHealthTimeout = 10
	DefaultThink         = false
)

type Config struct {
	TelegramToken string
	LLMProvider   string

	DatabaseDriver string
	DatabaseURL    string

	Ollama OllamaConfig
}

type OllamaConfig struct {
	BaseURL       string
	Model         string
	NumCtx        int
	NumPredict    int
	Temperature   float64
	HealthTimeout int
	Think         bool
}

func Load() Config {
	return Config{
		TelegramToken:  getEnv("TELEGRAM_TOKEN", ""),
		LLMProvider:    getEnv("LLM_PROVIDER", "ollama"),
		DatabaseDriver: getEnv("DATABASE_DRIVER", "sqlite"),
		DatabaseURL:    getEnv("DATABASE_URL", "file:storage.db"),

		Ollama: OllamaConfig{
			BaseURL:       getEnv("OLLAMA_BASE_URL", DefaultOllamaURL),
			Model:         getEnv("OLLAMA_MODEL", DefaultOllamaModel),
			NumCtx:        getIntEnv("OLLAMA_NUM_CTX", DefaultNumCtx),
			NumPredict:    getIntEnv("OLLAMA_NUM_PREDICT", DefaultNumPredict),
			Temperature:   getFloatEnv("OLLAMA_TEMPERATURE", DefaultTemperature),
			HealthTimeout: getIntEnv("OLLAMA_HEALTH_TIMEOUT", DefaultHealthTimeout),
			Think:         getBoolEnv("OLLAMA_THINK", DefaultThink),
		},
	}
}
