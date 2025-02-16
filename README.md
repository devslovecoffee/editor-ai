# Editor AI: AI-Powered Content Editing for Everyone

🚀 Automate text optimization with Large Language Models (LLMs)
Made by [@devslovecoffee](https://www.devslovecoffee.com/)

## Overview

When writing, it's crucial to focus on content creation without distractions. **Editor AI** is an open-source **AI-powered content editing tool** that automatically improves content using **Large Language Models (LLMs)**.

By leveraging **automated editing and SEO optimization**, Editor AI ensures your content remains **high-quality, engaging, and up-to-date**, without extra effort.

Why use Editor AI?
- AI-driven editing – Automatically refines your writing. 
- Boosts SEO – Keeps content fresh and optimized for search engines.
- Hands-free workflow – No manual editing required.

## How It Works

Editor AI scans and enhances your content using AI-based text optimization:
1. Reads content from the specified directory.
2. Uses **LLM-powered optimization** to refine and improve the text.
3. Applies the optimized changes to your files.
4. Saves the improved content automatically.

## Prerequisites

- Go installed (>=1.18)
- An OpenAI API key

## Installation

Clone the repository and navigate to the project directory:

```sh
git clone https://github.com/petttr1/editor-ai.git
cd editor-ai
```

## Usage

Run the program with the required flags:

```sh
go run main.go --dir "absolute/path/to/files" --api_key "your_openai_api_key"
```

Optionally, you can specify a glob pattern to filter specific files (e.g. markdown):

```sh
go run main.go --dir "absolute/path/to/files" --api_key "your_openai_api_key" --glob "*.md"
```

### Command-line Arguments

| Flag        | Description                                     | Required |
| ----------- | ----------------------------------------------- | -------- |
| `--dir`     | Absolute path to the directory containing files | ✅        |
| `--api_key` | OpenAI API key                                  | ✅        |
| `--glob`    | Glob pattern to filter files (default: `**`)    | ❌        |

### Customization

> **Note:** Currently, customization is only possible in the code. Future updates will include configuration file support. 

In order to achieve better results, you can customize the following:

| Variable    | Filepath                   | Note                                                                                                                                                   |
|-------------|----------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|
| `optimizeSystemPrompt`     | `internal/llm/prompts.go`  | Keep the `{{ .OutputRules }}` and `{{ .OutputFormat }}` within the prompt. Without this, the output may not be correct.                                |
| `optimizeUserPrompt` | `internal/llm/prompts.go`  | Keep the `{{ .OutputRules }}` and `{{ .OutputFormat }}` within the prompt. Without this, the output may not be correct.                                |
| `editRules`    | `internal/llm/rules.go`    | Edit if you have some custom rules you want the editor to incorporate.                                                                                 |
| `contentRules`    | `internal/llm/rules.go`    | Edit if you have specific content that you don't want to be edited (e.g. Code blocks, figures, configs).                                               |
| `replaceExamples`    | `internal/llm/examples.go` | If you notice the editor making a mistake on some part of your content, provide it as an example with the incorrect and the correct (expected output). |


## Example Output

```
Optimized 5 changes for article example.md
Optimized 3 changes for article another-file.md
```

## Roadmap
- [ ] Add release binaries.
- [ ] Parallel processing of content.
- [ ] Tests.
- [ ] Support for more LLM providers.
- [ ] LLM Model selection.
- [ ] Customizable prompt via config file.
- [ ] Customizable edit rules via config file.
- [ ] Customizable examples via config file.
- [ ] Store memories for later reference after optimizing.
- [ ] Read your content from the web.
- [ ] Optimization as part of your github workflow (optimize on PR).
- [ ] Add contributing guidelines.

## Contributing

🚀 We welcome contributions!

- Report issues or suggest features via GitHub Issues.
- Fork the repo and submit a Pull Request.
- Star ⭐ the repository if you find it useful!

Coming Soon: CONTRIBUTING.md file with setup instructions.

## Acknowledgements

The current editing rules are inspired by [Eva Parish's guide](https://evaparish.com/blog/how-i-edit), a fantastic resource on writing improvement.

## License

This project is licensed under the Apache 2.0 License.

