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
