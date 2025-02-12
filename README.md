# Editor AI

Made by [@devslovecoffee](https://www.devslovecoffee.com/)

## Overview

When writing, I find it best to focus solely on creating the content, without being distracted by editing. Automated content editing and optimization is now possible thanks to Large Language Models (LLMs). This tool was created to streamline the editing process by automatically refining and improving text files with your content. 

Additionally, periodic content updates help boost SEO and overall visibility, making Editor AI valuable for content creators who want to maintain high-quality, fresh content without the extra effort.

## How It Works

1. Content is read from the specified directory.
2. LLM is used to generate optimizations for the content.
3. Changes are applied to the files.
4. The optimized content is saved to the file.

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

> Note: For the time being, customization is possible solely in the code. Customizing via config files is on the roadmap. 

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

Feel free to submit issues or pull requests. No special rules for now.

## Acknowledgements

The current edit rules are based on this great article by [Eva Parish](https://evaparish.com/blog/how-i-edit).

## License

This project is licensed under the Apache 2.0 License.

