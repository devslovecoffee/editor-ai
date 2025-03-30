package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config represents the application configuration
type Config struct {
	Prompts  PromptsConfig  `yaml:"prompts"`
	Rules    RulesConfig    `yaml:"rules"`
	Examples ExamplesConfig `yaml:"examples"`
	Model    string         `yaml:"model"`
}

// PromptsConfig holds the customizable prompt templates
type PromptsConfig struct {
	SystemPrompt string `yaml:"system_prompt"`
	UserPrompt   string `yaml:"user_prompt"`
	OutputFormat string `yaml:"output_format"`
}

// RulesConfig holds the customizable rules
type RulesConfig struct {
	EditRules    string `yaml:"edit_rules"`
	ContentRules string `yaml:"content_rules"`
	OutputRules  string `yaml:"output_rules"`
}

// ExamplesConfig holds the customizable examples
type ExamplesConfig struct {
	ReplaceExamples string `yaml:"replace_examples"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		Prompts:  DefaultPrompts(),
		Rules:    DefaultRules(),
		Examples: DefaultExamples(),
		Model:    "gpt-4o-2024-08-06", // Default model
	}
}

// DefaultPrompts returns the default prompts configuration
func DefaultPrompts() PromptsConfig {
	return PromptsConfig{
		SystemPrompt: `Act as an expert writing coach and social media marketer.
Your job is to edit a piece of content following the edit rules provided to you.
Your main goal is to improve the writing and style.
Your secondary goal is to improve the content's visibility on search engines (SEO) to increase organic traffic.

When editing the content, ALWAYS follow these steps:
1. Carefully read and analyze the content.
2. Write a few sentences mentioning what is the main theme of the content, including the target audience, the kind of narrator, tone, mood, and key areas of improvements both in writing and SEO.
3. Carefully read the content again and identify opportunities to optimize the content. While doing so, follow the analysis you have written in the previous step, and the EDIT RULES.
4. Output the changes for the optimized content.
{{ if .EditRules }}

{{ .EditRules }}
{{ end }}
{{ if .ContentRules }}

{{ .ContentRules }}
{{ end }}

{{ .OutputRules }}

{{ .OutputFormat }}
{{ if .ReplaceExamples }}

{{ .ReplaceExamples }}
{{ end }}`,
		UserPrompt: `Improve the writing and style of the provided content.
Additionally, improve the content's visibility on search engines (SEO) to increase organic traffic.

Here is the content you are optimizing:
{{ .Content }}
{{ if or .EditRules .ContentRules }}

Remember the rules you need to follow:
{{ if .EditRules }}

{{ .EditRules }}
{{ end }}
{{ if .ContentRules }}

{{ .ContentRules }}
{{ end }}
{{ end }}

{{ .OutputRules }}

{{ .OutputFormat }}`,
		OutputFormat: `Output the changes for the optimized content in the following format:
<change>
	<search>[search term]</search>
	<replace>[replace term]</replace>
</change>`,
	}
}

// DefaultRules returns the default rules configuration
func DefaultRules() RulesConfig {
	return RulesConfig{
		EditRules: `Here are the rules that describe how to approach the editing process:
- Decide what you're actually saying.
- Who are you writing for?
- What is your main point? 

Repeat yourself (within reason)
- Look for ways that you can restate your point, clarify, or provide closure for the reader.

Simplify
- Strive to get to the point as quickly as possible.

Eliminate passive voice
- Rewrite a passive construction to active to make what you're saying clearer and make the sentence easier to read.

Don't use adverbs
- Replace an adverb with a better, more specific verb, or describe what you mean instead.

Don't assume knowledge
- Spell out acronyms on first use.

Be aware of your tone
- Know what kind of tone you're going for, and be consistent.
- Don't switch between formal / non-formal tones.`,
		ContentRules: `Here are the rules that describe how to handle the content:
- NEVER change the contents of a code block.
- NEVER change links.
- Avoid introducing any errors to the content structure (e.g. Yaml header).`,
		OutputRules: `Here are the rules that describe how to output the changes:
- For each optimization you find, output a separate change that can be replaced directly in the text.
- Each search term MUST be unique and match exactly once.
- Always consider the context of the full sentence before making a change and do not introduce grammatical errors.`,
	}
}

// DefaultExamples returns the default examples configuration
func DefaultExamples() ExamplesConfig {
	return ExamplesConfig{
		ReplaceExamples: `Here are some incorrect replacement examples that introduce various errors.
Use these only as examples, do not follow them word by word if you encounter a similar issue.

Examples:

Replace: The unfinished stuff from
Incorrect: Unfinished tasks from from
Correct: Unfinished tasks from

Replace: It works in the current, very simple, setup, but adding more imports
Incorrect: While this setup works, but adding more imports
Correct: While this setup works, adding more imports

Replace: we never have a whole side of the cube highlighted
Incorrect: we  the cube sides have inconsistent lighting coverage
Correct: the cube sides have inconsistent lighting coverage

Replace: Borrowing some code from this
Incorrect: Borrowing use the code example found at this
Correct: Borrowing the code example found at this

Replace: it's the new trend all the cool kids are doing
Incorrect: it's a trending design style popular among cool kids are doing
Correct: it's a trending design style popular among cool kids

Replace: Bento layouts actually don't tilt me, sorry
Incorrect: Bento layouts themselves don't tilt, apologies for the clickbait
Correct: Bento layouts themselves don't tilt me, apologies for the clickbait

Replace: tile is actually pretty simple and a really nice person - some Armando Canals already did that
Incorrect: tile is straightforward thanks to Armando Canals already did that
Correct: tile is straightforward thanks to Armando Canals who already did that

Replace: an extensible imports
Incorrect: an flexible imports
Correct: a flexible imports

Replace: Rotating then is a simple act of rotating layers
Incorrect: Rotating is simply of rotating layers
Correct: Rotating is simply the act of rotating layers

Replace: already did that
Incorrect: who has already did that
Correct: who has already done that`,
	}
}

// LoadConfig loads the configuration from the given file path
// If the file doesn't exist, it returns the default configuration
func LoadConfig(configPath string) (*Config, error) {
	// Start with default config
	config := DefaultConfig()

	// If no config file path is provided, return defaults
	if configPath == "" {
		return config, nil
	}

	// Check if file exists
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		return config, nil
	}

	// Read config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Parse YAML
	userConfig := new(Config)
	if err := yaml.Unmarshal(data, userConfig); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// Merge with defaults (only override non-empty fields)
	mergeConfig(config, userConfig)

	return config, nil
}

// FindConfigFile looks for a config file in standard locations
func FindConfigFile() string {
	// Look in current directory
	if _, err := os.Stat("editor-ai.yaml"); err == nil {
		return "editor-ai.yaml"
	}
	if _, err := os.Stat("editor-ai.yml"); err == nil {
		return "editor-ai.yml"
	}

	// Look in user's home directory
	home, err := os.UserHomeDir()
	if err == nil {
		configPath := filepath.Join(home, ".editor-ai.yaml")
		if _, err := os.Stat(configPath); err == nil {
			return configPath
		}
		configPath = filepath.Join(home, ".editor-ai.yml")
		if _, err := os.Stat(configPath); err == nil {
			return configPath
		}

		// Check .config directory
		configPath = filepath.Join(home, ".config", "editor-ai", "config.yaml")
		if _, err := os.Stat(configPath); err == nil {
			return configPath
		}
		configPath = filepath.Join(home, ".config", "editor-ai", "config.yml")
		if _, err := os.Stat(configPath); err == nil {
			return configPath
		}
	}

	return ""
}

// SaveDefaultConfig saves the default configuration as a template
func SaveDefaultConfig(path string) error {
	config := DefaultConfig()

	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	// Create directory if it doesn't exist
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// mergeConfig merges the user config into the default config
func mergeConfig(dest, src *Config) {
	// Prompts
	if src.Prompts.SystemPrompt != "" {
		dest.Prompts.SystemPrompt = src.Prompts.SystemPrompt
	}
	if src.Prompts.UserPrompt != "" {
		dest.Prompts.UserPrompt = src.Prompts.UserPrompt
	}
	if src.Prompts.OutputFormat != "" {
		dest.Prompts.OutputFormat = src.Prompts.OutputFormat
	}

	// Rules
	if src.Rules.EditRules != "" {
		dest.Rules.EditRules = src.Rules.EditRules
	}
	if src.Rules.ContentRules != "" {
		dest.Rules.ContentRules = src.Rules.ContentRules
	}
	if src.Rules.OutputRules != "" {
		dest.Rules.OutputRules = src.Rules.OutputRules
	}

	// Examples
	if src.Examples.ReplaceExamples != "" {
		dest.Examples.ReplaceExamples = src.Examples.ReplaceExamples
	}

	// Model
	if src.Model != "" {
		dest.Model = src.Model
	}
}
