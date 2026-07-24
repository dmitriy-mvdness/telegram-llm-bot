package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/config"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/model"
)

type OllamaClient struct {
	client *http.Client

	baseURL    string
	model      string
	numCtx     int
	numPredict int
	think      bool
}

func NewOllamaClient(cfg config.OllamaConfig) *OllamaClient {
	return &OllamaClient{
		client: &http.Client{
			Timeout: time.Duration(cfg.HealthTimeout) * time.Second,
		},
		baseURL:    cfg.BaseURL,
		model:      cfg.Model,
		numCtx:     cfg.NumCtx,
		numPredict: cfg.NumPredict,
		think:      cfg.Think,
	}
}

func (o *OllamaClient) Chat(ctx context.Context, messages []model.Message, options config.LLMOptions) (string, error) {
	reqBody := map[string]any{
		"model":    o.model,
		"messages": messages,
		"stream":   false,
		"think":    o.think,

		"options": map[string]any{
			"num_ctx":     o.numCtx,
			"num_predict": o.numPredict,
			"temperature": options.Temperature,
		},
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		o.baseURL+"/api/chat",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := o.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.Message.Content, nil
}
