package file

import (
	"fmt"
	"io/fs"
	"os"
	"path"

	"github.com/petttr1/editor-ai/internal/article"
)

type Reader struct{}

func NewReader() *Reader {
	return &Reader{}
}

func (r *Reader) Load(url, glob string) ([]*article.Article, error) {
	filepaths, err := r.getFilepaths(url, glob)
	if err != nil {
		return nil, fmt.Errorf("failed to get filepaths: %w", err)
	}

	articles := make([]*article.Article, 0)
	for _, v := range filepaths {
		content, err := r.readFileContent(v)
		if err != nil {
			fmt.Printf("failed to read file content: %v\n", err)
			continue
		}
		articles = append(articles, article.NewArticle(v, content))
	}

	return articles, nil
}

func (r *Reader) getFilepaths(url, glob string) ([]string, error) {
	root := os.DirFS(url)

	filteredFiles, err := fs.Glob(root, glob)

	if err != nil {
		return nil, fmt.Errorf("failed to get files: %w", err)
	}

	files := make([]string, 0)
	for _, v := range filteredFiles {
		files = append(files, path.Join(url, v))
	}

	return files, nil
}

func (r *Reader) readFileContent(filepath string) (string, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return string(file), nil
}
