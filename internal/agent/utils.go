package agent

import (
	"bytes"
	"fmt"
	"regexp"
	"text/template"
)

func ExtractTag(tag, content string) string {
	r := regexp.MustCompile(fmt.Sprintf(`<%s>((.|\n)*?)<\/%s>`, tag, tag))
	match := r.FindStringSubmatch(content)
	if len(match) > 1 {
		return match[1]
	}

	return ""
}

func ExtractTags(tag, content string) []string {
	r := regexp.MustCompile(fmt.Sprintf(`<%s>((.|\n)*?)<\/%s>`, tag, tag))
	matches := r.FindAllStringSubmatch(content, -1)
	if len(matches) <= 0 {
		return nil
	}

	matchedContent := make([]string, 0)
	for _, match := range matches {
		if len(match) > 1 {
			matchedContent = append(matchedContent, match[1])
		}
	}

	return matchedContent
}

func InsertValues(text string, val any) (string, error) {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("template").Parse(text)
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(buf, val)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
