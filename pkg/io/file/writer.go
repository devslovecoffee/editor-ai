package file

import (
	"fmt"
	"os"

	"github.com/petttr1/editor-ai/internal/article"
)

type Writer struct{}

func NewWriter() *Writer {
	return &Writer{}
}

func (w *Writer) Write(articles []*article.Article) error {
	for _, v := range articles {
		err := w.writeFile(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (w *Writer) writeFile(article *article.Article) error {
	err := os.WriteFile(article.FilePath, []byte(article.Content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
