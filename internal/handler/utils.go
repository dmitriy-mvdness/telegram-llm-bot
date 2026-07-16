package handler

func (h *Handler) isCommand(text string) bool {
	return len(text) > 0 && text[0] == '/'
}
