package llm

import (
	"reflect"
	"testing"
)

func TestExtractTag(t *testing.T) {
	tests := []struct {
		name     string
		tag      string
		content  string
		expected string
	}{
		{
			name:     "simple tag",
			tag:      "test",
			content:  "<test>content</test>",
			expected: "content",
		},
		{
			name:     "tag with newlines",
			tag:      "multi",
			content:  "<multi>line 1\nline 2</multi>",
			expected: "line 1\nline 2",
		},
		{
			name:     "tag with attributes in content",
			tag:      "data",
			content:  "<data>content with <attr>nested</attr> elements</data>",
			expected: "content with <attr>nested</attr> elements",
		},
		{
			name:     "nonexistent tag",
			tag:      "missing",
			content:  "<other>content</other>",
			expected: "",
		},
		{
			name:     "incomplete tag",
			tag:      "incomplete",
			content:  "<incomplete>content",
			expected: "",
		},
		{
			name:     "case sensitivity",
			tag:      "Case",
			content:  "<Case>Sensitive</Case>",
			expected: "Sensitive",
		},
		{
			name:     "multiple tags - returns first",
			tag:      "tag",
			content:  "<tag>first</tag>content<tag>second</tag>",
			expected: "first",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ExtractTag(test.tag, test.content)
			if result != test.expected {
				t.Errorf("Expected %q, got %q", test.expected, result)
			}
		})
	}
}

func TestExtractTags(t *testing.T) {
	tests := []struct {
		name     string
		tag      string
		content  string
		expected []string
	}{
		{
			name:     "multiple tags",
			tag:      "item",
			content:  "<item>first</item><item>second</item><item>third</item>",
			expected: []string{"first", "second", "third"},
		},
		{
			name:     "single tag",
			tag:      "solo",
			content:  "<solo>only one</solo>",
			expected: []string{"only one"},
		},
		{
			name:     "no tags",
			tag:      "missing",
			content:  "<other>content</other>",
			expected: nil,
		},
		{
			name:     "mixed content",
			tag:      "mix",
			content:  "text before <mix>content 1</mix> text between <mix>content 2</mix> text after",
			expected: []string{"content 1", "content 2"},
		},
		{
			name:     "tags with newlines",
			tag:      "multi",
			content:  "<multi>line 1\nline 2</multi>\n<multi>line 3\nline 4</multi>",
			expected: []string{"line 1\nline 2", "line 3\nline 4"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ExtractTags(test.tag, test.content)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestInsertValues(t *testing.T) {
	tests := []struct {
		name      string
		template  string
		values    interface{}
		expected  string
		expectErr bool
	}{
		{
			name:      "simple template",
			template:  "Hello, {{.Name}}!",
			values:    struct{ Name string }{"World"},
			expected:  "Hello, World!",
			expectErr: false,
		},
		{
			name:     "multiple values",
			template: "{{.Count}} {{.Item}}s",
			values: struct {
				Count int
				Item  string
			}{5, "apple"},
			expected:  "5 apples",
			expectErr: false,
		},
		{
			name:     "nested values",
			template: "{{.Person.Name}} is {{.Person.Age}} years old",
			values: struct {
				Person struct {
					Name string
					Age  int
				}
			}{struct {
				Name string
				Age  int
			}{"Alice", 30}},
			expected:  "Alice is 30 years old",
			expectErr: false,
		},
		{
			name:      "invalid template",
			template:  "{{.Name",
			values:    struct{ Name string }{"World"},
			expected:  "",
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := InsertValues(test.template, test.values)

			if test.expectErr {
				if err == nil {
					t.Error("Expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result != test.expected {
				t.Errorf("Expected %q, got %q", test.expected, result)
			}
		})
	}
}
