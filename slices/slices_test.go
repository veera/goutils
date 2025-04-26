package slices

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnique(t *testing.T) {
	got := Unique([]int{1, 2, 2, 3, 4, 4, 4, 5})
	want := []int{1, 2, 3, 4, 5}
	require.Equal(t, want, got)
}

func TestChunk(t *testing.T) {
	got := Chunk([]int{1, 2, 3, 4, 5}, 2)
	want := [][]int{{1, 2}, {3, 4}, {5}}
	require.Equal(t, want, got)
}

func TestContains(t *testing.T) {
	require.True(t, Contains([]int{1, 2, 3, 4}, 3))
	require.False(t, Contains([]int{1, 2, 3, 4}, 5))
}
