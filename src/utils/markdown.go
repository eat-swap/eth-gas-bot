package utils

import "strings"

const (
	badMarkdownChars = "_*`[\\]()#+-=|{}<>~."
)

func WrapForMarkdown(s string) string {
	for _, c := range badMarkdownChars {
		s = strings.ReplaceAll(s, string(c), "\\"+string(c))
	}
	return s
}
