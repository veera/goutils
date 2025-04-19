package strings

import (
	"regexp"
	"strings"
)

func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToUpper(s[:1]) + s[1:]
}

func KebabCase(s string) string {
	re := regexp.MustCompile(`[\s_]+`)
	s = re.ReplaceAllString(strings.TrimSpace(s), "-")
	return strings.ToLower(s)
}

func SnakeCase(s string) string {
	re := regexp.MustCompile(`[\s-]+`)
	s = re.ReplaceAllString(strings.TrimSpace(s), "_")
	return strings.ToLower(s)
}
