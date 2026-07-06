package main

import (
	"fmt"

	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/handler"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/service"
)

func main() {
	svc := service.New()
	h := handler.New(svc)

	result := h.Handle("hello")

	fmt.Println(result)
}
