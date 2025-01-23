package io

import "optimiseo/internal/article"

type Reader interface {
	Load(url string) ([]*article.Article, error)
}
