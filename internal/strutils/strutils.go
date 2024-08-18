package strutils

import (
	"strings"
	"unicode"
)


func StartsWith(str string, prefix string) bool {
    if len(str) < len(prefix) {
        return false
    }
    return str[:len(prefix)] == prefix
}

func IsWhitespace(s string) bool {
    return strings.TrimSpace(s) == ""
}

func RemoveDuplicates(input []string) []string {
    uniqueMap := make(map[string]bool)
    result := make([]string, 0, len(input))

    for _, s := range input {
        if !uniqueMap[s] {
            uniqueMap[s] = true
            result = append(result, s)
        }
    }

    return result
}

func Capitalize(str string) string {
    runes := []rune(str)
    if len(runes) == 0 {
        return ""
    }

    runes[0] = unicode.ToUpper(runes[0])

    for i := 1; i < len(runes); i++ {
        if unicode.IsSpace(runes[i-1]) {
            runes[i] = unicode.ToUpper(runes[i])
        }
    }

    return string(runes)
}

