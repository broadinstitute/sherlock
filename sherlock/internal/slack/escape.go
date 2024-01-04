package slack

import "strings"

func EscapeText(text string) string {
	replacements := map[string]string{
		"&": "&amp;",
		"<": "&lt;",
		">": "&gt;",
		"*": "\\*",
		"_": "\\_",
		"~": "\\~",
		"`": "\\`",
		"[": "\\[",
		"]": "\\]",
		"(": "\\(",
		")": "\\)",
	}
	for k, v := range replacements {
		text = strings.ReplaceAll(text, k, v)
	}
	return text
}
