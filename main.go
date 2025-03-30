package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/petttr1/editor-ai/internal/config"
	"github.com/petttr1/editor-ai/internal/llm"
	"github.com/petttr1/editor-ai/pkg/io/file"
)

func main() {
	// Define command-line flags
	dir := flag.String("dir", "", "absolute path to the directory with your content")
	apiKey := flag.String("api_key", "", "your openAI API key")
	globPattern := flag.String("glob", "**", "glob pattern to further filter the loaded files, optional")
	configPath := flag.String("config", "", "path to custom configuration file")
	initConfig := flag.Bool("init-config", false, "initialize a default configuration file")
	initConfigPath := flag.String("init-config-path", "", "path where to save the default configuration file")
	flag.Parse()

	// Check for init-config flag
	if *initConfig || *initConfigPath != "" {
		path := *initConfigPath
		if path == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				fmt.Println("Error getting home directory:", err)
				os.Exit(1)
			}
			path = filepath.Join(home, ".config", "editor-ai", "config.yaml")
		}

		if err := config.SaveDefaultConfig(path); err != nil {
			fmt.Printf("Failed to save default config: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Default configuration saved to %s\n", path)
		return
	}

	ctx := context.Background()

	// Check required parameters only if we're not just initializing a config
	if *dir == "" {
		fmt.Println("Error: dir is required")
		fmt.Println("Usage: editor-ai --dir <path> --api_key <key> [--glob <pattern>] [--config <path>]")
		os.Exit(1)
	}
	if *apiKey == "" {
		fmt.Println("Error: api_key is required")
		fmt.Println("Usage: editor-ai --dir <path> --api_key <key> [--glob <pattern>] [--config <path>]")
		os.Exit(1)
	}

	// Load config (from specified path or default locations)
	var cfg *config.Config
	var err error

	if *configPath != "" {
		cfg, err = config.LoadConfig(*configPath)
	} else {
		// Try to find config in standard locations
		foundConfig := config.FindConfigFile()
		cfg, err = config.LoadConfig(foundConfig)
	}

	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		os.Exit(1)
	}

	articles, err := file.NewReader().Load(*dir, *globPattern)
	if err != nil {
		fmt.Printf("Failed to load files: %v\n", err)
		os.Exit(1)
	}

	llmClient := llm.NewClientWithConfig(*apiKey, cfg)
	for _, article := range articles {
		changes, err := llmClient.GetOptimizedChanges(ctx, article)
		if err != nil {
			fmt.Printf("Failed to get optimized changes: %v\n", err)
			continue
		}

		article.AddChange(changes...)
		fmt.Printf("Optimized %d changes for article %s\n", len(changes), filepath.Base(article.FilePath))
		article.ApplyChanges()
	}

	writer := file.NewWriter()
	err = writer.Write(articles)
	if err != nil {
		fmt.Printf("Failed to write files: %v\n", err)
		os.Exit(1)
	}
}
