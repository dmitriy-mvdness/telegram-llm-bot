package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (o *OllamaClient) Health(ctx context.Context) error {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		o.baseURL+"/api/version",
		nil,
	)
	if err != nil {
		return err
	}

	resp, err := o.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	return nil
}

func (o *OllamaClient) Ready(ctx context.Context) error {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		o.baseURL+"/api/tags",
		nil,
	)
	if err != nil {
		return err
	}

	resp, err := o.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result struct {
		Models []struct {
			Name string `json:"name"`
		} `json:"models"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}

	for _, model := range result.Models {
		if model.Name == o.model {
			return nil
		}
	}

	return fmt.Errorf("ollama model %s not found", o.model)
}
