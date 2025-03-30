package article

import (
	"testing"
)

func TestApplyChanges(t *testing.T) {
	article := NewArticle("test.txt", "Hello [World]")
	change := &Change{Search: "[World]", Replace: "[Universe]"}
	article.AddChange(change)
	article.ApplyChanges()

	if article.Content != "Hello [Universe]" {
		t.Errorf("Expected 'Hello [Universe]', got '%s'", article.Content)
	}

	// Test with no changes
	article = NewArticle("test.txt", "Hello World")
	article.ApplyChanges()

	if article.Content != "Hello World" {
		t.Errorf("Expected 'Hello World', got '%s'", article.Content)
	}

	// Test with multiple occurrences
	article = NewArticle("test.txt", "Hello [World] [World]")
	article.AddChange(change)
	article.ApplyChanges()

	if article.Content != "Hello [World] [World]" {
		t.Errorf("Expected 'Hello [World] [World]', got '%s'", article.Content)
	}
}
