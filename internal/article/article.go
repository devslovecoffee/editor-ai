package article

import "strings"

type Article struct {
	FilePath string
	Content  string
	Changes  []*Change
}

func NewArticle(filePath, content string) *Article {
	return &Article{
		FilePath: filePath,
		Content:  content,
		Changes:  make([]*Change, 0),
	}
}

func (a *Article) AddChange(change ...*Change) {
	a.Changes = append(a.Changes, change...)
}

func (a *Article) ApplyChanges() {
	for _, change := range a.Changes {
		count := strings.Count(a.Content, change.Search)
		if count == 0 || count > 1 {
			continue
		}

		searchSanitized := strings.TrimPrefix(change.Search, "[")
		searchSanitized = strings.TrimSuffix(change.Search, "]")

		replaceSanitized := strings.TrimPrefix(change.Replace, "[")
		replaceSanitized = strings.TrimSuffix(change.Replace, "]")

		a.Content = strings.Replace(a.Content, searchSanitized, replaceSanitized, 1)
	}
}
