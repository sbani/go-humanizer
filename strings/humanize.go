package strings

import (
	"regexp"
	"strings"
)

// ForbiddenWords are words that will be removed from text (in Humanize func)
var ForbiddenWords = []string{
	"id",
}

// Humanize a string means to remove underscores and return string as lower.
// String can be capitalized
func Humanize(s string, capitalize bool) string {
	s = regexp.MustCompile("([A-Z])").ReplaceAllString(s, "_$1")
	s = regexp.MustCompile("[_\\s]+").ReplaceAllString(s, " ")
	s = strings.ToLower(s)

	for _, word := range ForbiddenWords {
		s = strings.Replace(s, word, "", -1)
	}

	s = strings.TrimSpace(s)

	if capitalize {
		s = strings.ToUpper(s[:1]) + s[1:]
	}

	return s
}
