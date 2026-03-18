package providers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/zakirkun/ice-tea/internal/analyzer/llm"
)

// OpenAIProvider implements the llm.Provider interface for OpenAI API
type OpenAIProvider struct {
	apiKey string
	model  string
	client *http.Client
}

// NewOpenAIProvider creates a new OpenAI provider
func NewOpenAIProvider(apiKeyEnv, model string) (*OpenAIProvider, error) {
	apiKey := strings.TrimSpace(os.Getenv(apiKeyEnv))
	if apiKey == "" {
		return nil, fmt.Errorf("OpenAI API key not found in environment variable %s", apiKeyEnv)
	}

	return &OpenAIProvider{
		apiKey: apiKey,
		model:  model,
		client: &http.Client{},
	}, nil
}

type openAIRequest struct {
	Model          string          `json:"model"`
	Messages       []openAIMessage `json:"messages"`
	ResponseFormat map[string]string `json:"response_format"`
	Temperature    float32         `json:"temperature"`
}

type openAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openAIResponse struct {
	Choices []struct {
		Message openAIMessage `json:"message"`
	} `json:"choices"`
}

// Analyze sends the request to OpenAI
func (p *OpenAIProvider) Analyze(ctx context.Context, req llm.AnalysisRequest) (*llm.AnalysisResponse, error) {
	// 1. We must mock the prompt builder here, or ideally it should be exported
	// For this skeleton implementation, we will use a simplified approach:
	sysPrompt := "You are a security expert. Respond in JSON based on the provided code."
	userPrompt := fmt.Sprintf("Analyze: %s\nCode:\n%s", req.Rule.Description, req.CodeSnippet)

	body := openAIRequest{
		Model:       p.model,
		Temperature: 0.2, // Low temp for more deterministic analysis
		ResponseFormat: map[string]string{"type": "json_object"},
		Messages: []openAIMessage{
			{Role: "system", Content: sysPrompt},
			{Role: "user", Content: userPrompt},
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+p.apiKey)

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("OpenAI API error: %d %s", resp.StatusCode, string(respBody))
	}

	var openAIResp openAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&openAIResp); err != nil {
		return nil, err
	}

	if len(openAIResp.Choices) == 0 {
		return nil, fmt.Errorf("no response from OpenAI")
	}

	var analysis llm.AnalysisResponse
	if err := json.Unmarshal([]byte(openAIResp.Choices[0].Message.Content), &analysis); err != nil {
		return nil, fmt.Errorf("failed to parse JSON from LLM: %w", err)
	}

	return &analysis, nil
}
