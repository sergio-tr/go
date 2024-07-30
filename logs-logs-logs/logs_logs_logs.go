package logs

import (
	"strings"
	"unicode/utf8"
)

const (
	Recommendation = "recommendation"
	Search         = "search"
	Weather        = "weather"
	Default        = "default"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	mapper := map[string]rune{
		Recommendation: '❗',
		Search:         '🔍',
		Weather:        '☀',
	}

	for _, char := range log {

		for mapperString, mapperChar := range mapper {
			if mapperChar == char {
				return mapperString
			}
		}

	}

	return Default
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	builder := strings.Builder{}
	for _, char := range log {

		if char == oldRune {
			builder.WriteRune(newRune)
		} else {
			builder.WriteRune(char)
		}
	}

	return builder.String()
}

// WithinLimit determines whether the number of characters in log is within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
