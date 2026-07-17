package handler

import "github.com/go-telegram/bot/models"

func SettingsKeyboard() *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{
					Text:         "🗑️ Очистить историю",
					CallbackData: "clear_history",
				},
				{
					Text:         "🎭 Роль ассистента",
					CallbackData: "system_prompts",
				},
			},
		},
	}
}

func AssistantRoleKeyboard() *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{
					Text:         "🤖 Обычный",
					CallbackData: "prompt_default",
				},
			},
			{
				{
					Text:         "⚡ Краткий",
					CallbackData: "prompt_concise",
				},
			},
			{
				{
					Text:         "🎓 Экспертный",
					CallbackData: "prompt_academic",
				},
			},
			{
				{
					Text:         "🧠 Наводящий",
					CallbackData: "prompt_provocative",
				},
			},
			{
				{
					Text:         "📚 Подробный",
					CallbackData: "prompt_encyclopedic",
				},
			},
		},
	}
}
