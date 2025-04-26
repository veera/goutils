package times

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAddDays(t *testing.T) {
	base := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	got := AddDays(base, 5)
	require.Equal(t, base.AddDate(0, 0, 5), got)
}

func TestFormatDate(t *testing.T) {
	d := time.Date(2025, 4, 26, 15, 0, 0, 0, time.UTC)
	got := FormatDate(d)
	require.Equal(t, "2025-04-26", got)
}

func TestDaysBetween(t *testing.T) {
	a := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	b := time.Date(2025, 1, 5, 0, 0, 0, 0, time.UTC)
	got := DaysBetween(a, b)
	require.Equal(t, 4, got)
}
