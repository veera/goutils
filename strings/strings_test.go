package strings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCapitalize(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello", "Hello"},
		{"", ""},
		{"gO", "GO"},
	}

	for _, tt := range tests {
		require.Equal(t, tt.want, Capitalize(tt.input))
	}
}

func TestKebabCase(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"Hello World_test", "hello-world-test"},
		{"Already-Kebab", "already-kebab"},
		{" spaced out_text", "spaced-out-text"},
	}

	for _, tt := range tests {
		require.Equal(t, tt.want, KebabCase(tt.input))
	}
}

func TestSnakeCase(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"Hello-World test", "hello_world_test"},
		{"Multiple   spaces-here", "multiple_spaces_here"},
	}
	for _, tt := range tests {
		require.Equal(t, tt.want, SnakeCase(tt.input))
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello", "olleh"},
		{"Go", "oG"},
		{"", ""},
		{"a", "a"},
		{"racecar", "racecar"},
		{"123 💡 abc", "cba 💡 321"},
	}

	for _, tt := range tests {
		require.Equal(t, tt.want, Reverse(tt.input), "Reverse(%q)", tt.input)
	}
}
