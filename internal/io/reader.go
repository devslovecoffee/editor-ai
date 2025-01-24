package io

import "optimiseo/internal/article"

type Reader interface {
	Load(url string, glob string) ([]*article.Article, error)
}
