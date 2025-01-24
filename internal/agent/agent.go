package agent

import (
	"context"
	"fmt"
	article2 "optimiseo/internal/article"

	"github.com/sashabaranov/go-openai"
)

type Agent struct {
	client *openai.Client
}

func NewAgent(token string) *Agent {
	return &Agent{
		client: openai.NewClient(token),
	}
}

func (a *Agent) GetOptimizedChanges(ctx context.Context, article *article2.Article) (
	changes []*article2.Change,
	err error,
) {
	systemMessage, err := InsertValues(
		optimizeSystemPrompt, map[string]any{
			"EditRules":       &editRules,
			"ContentRules":    &contentRules,
			"OutputRules":     &outputRules,
			"OutputFormat":    &outputFormat,
			"ReplaceExamples": &replaceExamples,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert values: %w", err)
	}

	userMessage, err := InsertValues(
		optimizeUserPrompt, map[string]any{
			"Content":      article.Content,
			"EditRules":    &editRules,
			"ContentRules": &contentRules,
			"OutputRules":  &outputRules,
			"OutputFormat": &outputFormat,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert values: %w", err)
	}

	resp, err := a.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT4o20240806,
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
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return nil, fmt.Errorf("failed to get chat completion: %w", err)
	}

	responseContent := resp.Choices[0].Message.Content

	fmt.Printf("Response for article: %s\n", article.FilePath)
	fmt.Println(responseContent)

	responseChanges := ExtractTags("change", responseContent)
	if responseChanges == nil {
		return nil, fmt.Errorf("failed to extract changes")
	}

	changes = make([]*article2.Change, 0)
	for _, change := range responseChanges {
		search := ExtractTag("search", change)
		if search == "" {
			continue
		}

		replace := ExtractTag("replace", change)

		changes = append(
			changes, &article2.Change{
				Search:  search,
				Replace: replace,
			},
		)
	}
	return changes, nil
}
