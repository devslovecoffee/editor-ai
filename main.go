package main

import (
	"context"
	"flag"
	"fmt"
	"optimiseo/internal/agent"
	"optimiseo/pkg/io/file"
	"path/filepath"
)

func main() {
	dir := flag.String("dir", "", "absolute path to the directory with your content")
	apiKey := flag.String("api_key", "", "your openAI API key")
	globPattern := flag.String("glob", "**", "glob pattern to further filter the loaded files, optional")
	flag.Parse()

	ctx := context.Background()

	if *dir == "" {
		panic("dir is required")
	}
	if *apiKey == "" {
		panic("api key is required")
	}

	articles, err := file.NewReader().Load(*dir, *globPattern)
	if err != nil {
		panic(err)
	}

	llmAgent := agent.NewAgent(*apiKey)
	for _, article := range articles {
		changes, err := llmAgent.GetOptimizedChanges(ctx, article)
		if err != nil {
			fmt.Printf("failed to get optimized changes: %v\n", err)
			continue
		}
		article.AddChange(changes...)
		fmt.Printf("Optimized %d changes for article %s\n", len(changes), filepath.Base(article.FilePath))
		article.ApplyChanges()
	}

	writer := file.NewWriter()
	err = writer.Write(articles)
	if err != nil {
		panic(err)
	}
}
