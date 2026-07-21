package cmd

import (
	"github.com/spf13/viper"
	"io"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func makeRequest() (string, error) {
	systemPrompt, err := generateSystemPrompt()
	if err != nil {
		return "", err
	}

	prompt, err := generatePrompt()
	if err != nil {
		return "", err
	}

	body := map[string]any{
		"model":    viper.GetString("model"),
		"messages": []map[string]string{
			{"role": "system", "content": systemPrompt},
			{"role": "user", "content": prompt},
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("marshal error: %w", err)
	}

	req, err := http.NewRequest("POST", viper.GetString("base_url") + "/chat/completions", bytes.NewReader(jsonBody))
	if err != nil {
		return "", fmt.Errorf("request error: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + viper.GetString("api_key"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("request error: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
    	return "", fmt.Errorf("unexpected status %d: %s", resp.StatusCode, respBody)
	}

	return string(respBody), nil
}
