package strings

import "strings"

// Truncate a string, but never cut a word. So if you would cut the word it will
// be return completely (to the end or the next whitespace)
func Truncate(s string, charactersCount int) string {
	strLen := len(s)
	if charactersCount < 0 || strLen <= charactersCount {
		return s
	}

	length := strLen
	breakpoint := strings.Index(s[charactersCount:], " ")
	if breakpoint != -1 {
		length = breakpoint + charactersCount
	} else if s[(charactersCount-1):charactersCount] == " " {
		length = charactersCount
	}

	return strings.TrimRight(s[:length], " ")
}
