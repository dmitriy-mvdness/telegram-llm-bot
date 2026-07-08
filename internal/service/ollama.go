package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/config"
)

type OllamaClient struct {
	client *http.Client

	baseURL     string
	model       string
	numCtx      int
	numPredict  int
	temperature float64
	think       bool
}

func NewOllamaClient(cfg config.OllamaConfig) *OllamaClient {
	return &OllamaClient{
		client: &http.Client{
			Timeout: time.Duration(cfg.HealthTimeout) * time.Second,
		},
		baseURL:     cfg.BaseURL,
		model:       cfg.Model,
		numCtx:      cfg.NumCtx,
		numPredict:  cfg.NumPredict,
		temperature: cfg.Temperature,
		think:       cfg.Think,
	}
}

func (o *OllamaClient) Generate(prompt string) (string, error) {
	reqBody := map[string]any{
		"model":  o.model,
		"prompt": prompt,
		"stream": false,
		"think":  o.think,

		"options": map[string]any{
			"num_ctx":     o.numCtx,
			"num_predict": o.numPredict,
			"temperature": o.temperature,
		},
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(
		o.baseURL+"/api/generate",
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Response string `json:"response"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.Response, nil
}
