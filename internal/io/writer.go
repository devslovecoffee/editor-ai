package io

import "optimiseo/internal/article"

type Writer interface {
	Write(articles []*article.Article) error
}
