package io

import "github.com/petttr1/editor-ai/internal/article"

type Reader interface {
	Load(url string, glob string) ([]*article.Article, error)
}
