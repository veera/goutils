package times

import "time"

func Now() time.Time {
	return time.Now()
}

func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func DaysBetween(a, b time.Time) int {
	if a.After(b) {
		a, b = b, a
	}

	return int(b.Sub(a).Hours() / 24)
}
