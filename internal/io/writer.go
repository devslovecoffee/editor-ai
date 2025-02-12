package io

import "github.com/petttr1/editor-ai/internal/article"

type Writer interface {
	Write(articles []*article.Article) error
}
