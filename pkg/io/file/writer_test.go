package file

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/petttr1/editor-ai/internal/article"
)

func TestWrite(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "writer-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test articles
	testArticles := []*article.Article{
		article.NewArticle(filepath.Join(tempDir, "article1.md"), "# Test Article 1"),
		article.NewArticle(filepath.Join(tempDir, "article2.md"), "# Test Article 2"),
	}

	// Test writing articles
	writer := NewWriter()
	err = writer.Write(testArticles)
	if err != nil {
		t.Fatalf("Failed to write articles: %v", err)
	}

	// Verify content was written correctly
	for _, a := range testArticles {
		content, err := os.ReadFile(a.FilePath)
		if err != nil {
			t.Errorf("Failed to read written file %s: %v", a.FilePath, err)
			continue
		}

		if string(content) != a.Content {
			t.Errorf("Expected content %q, got %q for file %s", a.Content, string(content), a.FilePath)
		}
	}
}

func TestWriteFile(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "writefile-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Test writing a single article
	filePath := filepath.Join(tempDir, "test-article.md")
	testContent := "# Test Content for WriteFile"
	testArticle := article.NewArticle(filePath, testContent)

	writer := NewWriter()
	err = writer.writeFile(testArticle)
	if err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}

	// Verify content was written correctly
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read written file: %v", err)
	}

	if string(content) != testContent {
		t.Errorf("Expected content %q, got %q", testContent, string(content))
	}

	// Test writing to an invalid location
	invalidArticle := article.NewArticle("/nonexistent/directory/file.md", "Invalid content")
	err = writer.writeFile(invalidArticle)
	if err == nil {
		t.Error("Expected error when writing to invalid location, got nil")
	}
}
