package utils

import "strings"

const (
	badMarkdownChars   = "_*`[\\]()#+-=|{}<>~."
	worseMarkdownChars = "\\#+-=|{}<>~."
)

func WrapForMarkdown(s string) string {
	for _, c := range badMarkdownChars {
		s = strings.ReplaceAll(s, string(c), "\\"+string(c))
	}
	return s
}

func WrapForMarkdownWorse(s string) string {
	for _, c := range worseMarkdownChars {
		s = strings.ReplaceAll(s, string(c), "\\"+string(c))
	}
	return s
}
