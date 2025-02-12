package llm

import (
	"context"
	"fmt"

	article2 "github.com/petttr1/editor-ai/internal/article"

	"github.com/sashabaranov/go-openai"
)

type Client struct {
	client *openai.Client
}

func NewClient(token string) *Client {
	return &Client{
		client: openai.NewClient(token),
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
		optimizeSystemPrompt, map[string]any{
			"EditRules":       &editRules,
			"ContentRules":    &contentRules,
			"OutputRules":     outputRules,
			"OutputFormat":    outputFormat,
			"ReplaceExamples": &replaceExamples,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert system promtp values: %w", err)
	}

	userMessage, err := InsertValues(
		optimizeUserPrompt, map[string]any{
			"Content":      content,
			"EditRules":    &editRules,
			"ContentRules": &contentRules,
			"OutputRules":  outputRules,
			"OutputFormat": outputFormat,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user prompt values: %w", err)
	}

	return &openai.ChatCompletionRequest{
		Model: model,
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
