# Editor AI: AI-Powered Content Editing for Everyone
by [@devslovecoffee](https://www.devslovecoffee.com/)

<p align="center" width="100%">
    <img src="./assets/EAi_logo.svg" alt="Editor AI Logo" style="width: 20%;">
    <br><em>🚀 Automate text optimization with Large Language Models (LLMs)</em>
</p>

## Overview

Writing requires a focus on content creation without distractions. **Editor AI** is an open-source **AI-powered content editing tool** that automatically enhances content with **Large Language Models (LLMs)**.

By utilizing **automated editing and search engine optimization (SEO)**, Editor AI ensures your content remains **high-quality, engaging, and up-to-date**, without extra effort.

Benefits of Using Editor AI
- AI-driven editing – It refines your writing efficiently. 
- Boosts SEO – Keeps content fresh and optimized for search engines.
- Hands-free workflow – It eliminates the need for manual editing.

## How It Works

Editor AI improves your content through AI-based text optimization:
1. Reads content from the specified directory.
2. Uses **LLM-powered optimization** to refine and improve the text.
3. Applies any enhancements directly to your files.
4. Saves the improved content automatically.

## Prerequisites

- Go installed (>=1.18)
- An OpenAI API key

## Installation

To get started, you have two options:

### Option 1: Use pre-built binaries (recommended)

Download the latest release binary for your operating system from the [GitHub Releases page](https://github.com/petttr1/editor-ai/releases).

```sh
# Linux
chmod +x editor-ai-linux-amd64
./editor-ai-linux-amd64 --dir "/path/to/files" --api_key "your_openai_api_key"

# macOS
chmod +x editor-ai-darwin-amd64  # For Intel Macs
./editor-ai-darwin-amd64 --dir "/path/to/files" --api_key "your_openai_api_key"

# macOS (Apple Silicon)
chmod +x editor-ai-darwin-arm64
./editor-ai-darwin-arm64 --dir "/path/to/files" --api_key "your_openai_api_key"

# Windows
editor-ai-windows-amd64.exe --dir "C:\path\to\files" --api_key "your_openai_api_key"
```

### Option 2: Build from source

Clone the repository and navigate to the project directory:

```sh
git clone https://github.com/petttr1/editor-ai.git
cd editor-ai
```

Build using the provided Makefile:

```sh
# Build the application
make build

# Run tests
make test

# Build releases for all platforms
make release

# Run the application
make run DIR="/path/to/files" API_KEY="your_openai_api_key" GLOB="*.md"
```

Run `make help` to see all available commands.

## Usage

Run the program with the required flags:

```sh
go run main.go --dir "absolute/path/to/files" --api_key "your_openai_api_key"
```

If desired, specify a glob pattern to target particular files (e.g., markdown):

```sh
go run main.go --dir "absolute/path/to/files" --api_key "your_openai_api_key" --glob "*.md"
```

### Command-line Arguments

| Flag                | Description                                            | Required |
| ------------------- | ------------------------------------------------------ | -------- |
| `--dir`             | Absolute path to the directory containing files        | ✅        |
| `--api_key`         | OpenAI API key                                         | ✅        |
| `--glob`            | Glob pattern to filter files (default: `**`)           | ❌        |
| `--config`          | Path to custom configuration file                      | ❌        |
| `--init-config`     | Initialize a default configuration file                | ❌        |
| `--init-config-path`| Path where to save the default configuration file      | ❌        |

### Customization

You can customize OptimSEO's behavior using configuration files in YAML format. This allows you to adapt the tool to different content types and requirements without modifying the code.

#### Creating a Configuration File

To create a default configuration file:

```sh
# Create a config file in the default location (~/.config/editor-ai/config.yaml)
editor-ai --init-config

# Create a config file in a custom location
editor-ai --init-config-path ./my-config.yaml
```

#### Using a Configuration File

To use a configuration file:

```sh
# Use a specific config file
editor-ai --dir "/path/to/files" --api_key "your_openai_api_key" --config ./my-config.yaml

# Use the default config file locations (searches in current directory, home directory, etc.)
editor-ai --dir "/path/to/files" --api_key "your_openai_api_key"
```

The tool will automatically look for configuration files in these locations (in order):
1. Current directory (`editor-ai.yaml` or `editor-ai.yml`)
2. User's home directory (`.editor-ai.yaml` or `.editor-ai.yml`)
3. `~/.config/editor-ai/config.yaml` or `~/.config/editor-ai/config.yml`

#### Configuration File Structure

```yaml
# LLM Model to use
model: gpt-4o-2024-08-06

# Prompts Configuration
prompts:
  system_prompt: |
    # Your system prompt template
  user_prompt: |
    # Your user prompt template
  output_format: |
    # Your output format template

# Rules Configuration
rules:
  edit_rules: |
    # Your editing rules
  content_rules: |
    # Your content handling rules
  output_rules: |
    # Your output formatting rules

# Examples Configuration
examples:
  replace_examples: |
    # Your replacement examples
```

See the `examples/` directory for complete configuration examples:
- `examples/config.yaml` - Default configuration
- `examples/marketing-config.yaml` - Specialized for marketing content
- `examples/technical-docs-config.yaml` - Specialized for technical documentation

> **Note:** The configuration file uses Go templates with variables like `{{ .Content }}`. Make sure to preserve these variables in your custom prompts to ensure proper functionality.

## Example Output

```
Optimized 5 changes for article example.md
Optimized 3 changes for article another-file.md
```

## Roadmap
- [x] Add release binaries.
- [ ] Parallel processing of content.
- [x] Tests.
- [ ] Support for more LLM providers.
- [ ] LLM Model selection.
- [x] Customizable prompt via config file.
- [x] Customizable edit rules via config file.
- [x] Customizable examples via config file.
- [ ] Store memories for later reference after optimizing.
- [ ] Read your content from the web.
- [ ] Optimization as part of your github workflow (optimize on PR).
- [ ] Add contributing guidelines.

## Contributing

🚀 Contributions are welcome!

- Report issues or suggest features via GitHub Issues.
- Fork the repo and submit a Pull Request.
- Star ⭐ the repository if you find it useful!

Coming Soon: CONTRIBUTING.md file with setup instructions.

## Acknowledgements

The current editing rules are inspired by [Eva Parish's guide](https://evaparish.com/blog/how-i-edit), a fantastic resource on writing improvement.

## License

This project is licensed under the Apache 2.0 License.

