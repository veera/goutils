package maps

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	got := Keys(m)
	require.ElementsMatch(t, []string{"a", "b", "c"}, got)
}

func TestValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	got := Values(m)
	require.ElementsMatch(t, []int{1, 2, 3}, got)
}

func TestInvert(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2}
	got := Invert(m)

	require.Equal(t, map[int]string{1: "one", 2: "two"}, got)
}
