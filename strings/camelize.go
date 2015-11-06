package strings

import (
    "strings"
    "regexp"
)

// Camelize the string
func Camelize(s string) string {
    words := strings.Split(s, "_")
    for i, word := range words {
        words[i] = ucFirst(strings.ToLower(word))
    }
    
    return lcFirst(strings.Join(words, ""))
}

// Underscore the string
func Underscore(s string) string {
    reg := regexp.MustCompile("[A-Z]+")
    indexes := reg.FindAllStringIndex(s, -1)
    laststart := 0
    words := make([]string, len(indexes) + 1)
    for i, element := range indexes {
        words[i] = s[laststart:element[0]]
        element[1]--
        laststart = element[1]
    }
    words[len(indexes)] = s[laststart:]
    
    return strings.ToLower(strings.Join(words, "_"))
}

// Lower case first letter of given string
func lcFirst(s string) string {
    if len(s) <= 1 {
        return strings.ToLower(s)
    }
    return strings.ToLower(s[:1]) + s[1:]
}

// Upper case first letter of given string
func ucFirst(s string) string {
    if len(s) <= 1 {
        return strings.ToUpper(s)
    }
    return strings.ToUpper(s[:1]) + s[1:]
}
