package handler

import "fmt"

func (h *Handler) isCommand(text string) bool {
	return len(text) > 0 && text[0] == '/'
}

func PromptSelectedMessage(displayName string) string {
	return fmt.Sprintf(`✅ Выбран стиль: %s

Теперь я буду отвечать в выбранном режиме.
Вы можете изменить стиль в любой момент через настройки.`, displayName)
}
