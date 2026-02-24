package utils

import "strings"

// CleanInput removes surrounding whitespace and converts to lowercase
func CleanInput(input string) string {
	return strings.TrimSpace(strings.ToLower(input))
}
