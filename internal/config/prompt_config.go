package config

type PromptConfig struct {
	Options LLMOptions
}

var PromptConfigs = map[string]PromptConfig{
	"default": {
		Options: LLMOptions{
			Temperature: 0.5,
		},
	},

	"concise": {
		Options: LLMOptions{
			Temperature: 0.3,
		},
	},

	"academic": {
		Options: LLMOptions{
			Temperature: 0.2,
		},
	},

	"provocative": {
		Options: LLMOptions{
			Temperature: 0.6,
		},
	},

	"encyclopedic": {
		Options: LLMOptions{
			Temperature: 0.25,
		},
	},
}
