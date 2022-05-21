package util

import "strings"

func Title(text string) string {
	if len(text) < 2 {
		return text
	}
	return strings.ToUpper(text[0:1]) + text[1:]
}
