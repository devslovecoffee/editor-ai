package llm

import (
	"context"
	"fmt"

	article2 "github.com/petttr1/editor-ai/internal/article"
	"github.com/petttr1/editor-ai/internal/config"

	"github.com/sashabaranov/go-openai"
)

type Client struct {
	client *openai.Client
	config *config.Config
}

func NewClient(token string) *Client {
	return &Client{
		client: openai.NewClient(token),
		config: config.DefaultConfig(),
	}
}

func NewClientWithConfig(token string, cfg *config.Config) *Client {
	return &Client{
		client: openai.NewClient(token),
		config: cfg,
	}
}

func (c *Client) GetOptimizedChanges(ctx context.Context, article *article2.Article) (
	changes []*article2.Change,
	err error,
) {
	request, err := c.createRequest(article.Content)
	if err != nil {
		return nil, fmt.Errorf("failed to create chat completion request: %w", err)
	}

	resp, err := c.client.CreateChatCompletion(ctx, *request)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return nil, fmt.Errorf("failed to get chat completion: %w", err)
	}

	responseContent := resp.Choices[0].Message.Content

	// Uncomment for debugging the response content
	// fmt.Printf("Response for article: %s\n", article.FilePath)
	// fmt.Println(responseContent)

	return c.extractChanges(responseContent)
}

func (c *Client) createRequest(content string) (*openai.ChatCompletionRequest, error) {
	systemMessage, err := InsertValues(
		c.config.Prompts.SystemPrompt, map[string]any{
			"EditRules":       &c.config.Rules.EditRules,
			"ContentRules":    &c.config.Rules.ContentRules,
			"OutputRules":     c.config.Rules.OutputRules,
			"OutputFormat":    c.config.Prompts.OutputFormat,
			"ReplaceExamples": &c.config.Examples.ReplaceExamples,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert system promtp values: %w", err)
	}

	userMessage, err := InsertValues(
		c.config.Prompts.UserPrompt, map[string]any{
			"Content":      content,
			"EditRules":    &c.config.Rules.EditRules,
			"ContentRules": &c.config.Rules.ContentRules,
			"OutputRules":  c.config.Rules.OutputRules,
			"OutputFormat": c.config.Prompts.OutputFormat,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user prompt values: %w", err)
	}

	// Get model from config or use default
	modelName := getModelName(c.config.Model)

	return &openai.ChatCompletionRequest{
		Model: modelName,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemMessage,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: userMessage,
			},
		},
	}, nil
}

func (c *Client) extractChanges(content string) ([]*article2.Change, error) {
	changes := ExtractTags("change", content)
	if changes == nil {
		return nil, fmt.Errorf("failed to extract changes")
	}

	articleChanges := make([]*article2.Change, 0)
	for _, change := range changes {
		search := ExtractTag("search", change)
		if search == "" {
			continue
		}

		replace := ExtractTag("replace", change)

		articleChanges = append(
			articleChanges, &article2.Change{
				Search:  search,
				Replace: replace,
			},
		)
	}
	return articleChanges, nil
}

// getModelName maps a user-friendly model name to an OpenAI SDK model identifier
func getModelName(model string) string {
	// Map of user-friendly names to OpenAI SDK constants
	modelMap := map[string]string{
		"gpt-4o-2024-08-06": openai.GPT4o20240806,
		"gpt-4o":            openai.GPT4o,
		"gpt-4-turbo":       openai.GPT4Turbo,
		"gpt-3.5-turbo":     openai.GPT3Dot5Turbo,
	}

	// If model is a known key, return the mapped value
	if val, ok := modelMap[model]; ok {
		return val
	}

	// If model isn't in our map, assume it's a direct OpenAI model name
	// This allows users to specify exact model names if desired
	return model
}
