package time

import (
	"fmt"
	"math"
	"time"
)

// Difference finds the difference betwee t1 and t2.
func Difference(t1, t2 time.Time) string {
	seconds := math.Abs(t1.Sub(t2).Seconds())
	if seconds < 1 {
		return "just now"
	}
	var end string
	if t1.After(t2) {
		end = "ago"
	} else {
		end = "from now"
	}
	// Minute
	minutes := seconds / 60
	if minutes < 1 {
		return createString(seconds, "second", end)
	}
	// Hour
	hours := minutes / 60
	if hours < 1 {
		return createString(minutes, "minute", end)
	}
	// Days
	days := hours / 24
	if days < 1 {
		return createString(hours, "hour", end)
	}
	// Weeks
	weeks := days / 7
	if weeks < 1 {
		return createString(days, "day", end)
	}
	// Months
	months := weeks / 4
	if months < 1 {
		return createString(weeks, "week", end)
	}
	// Years
	years := months / 12
	if years < 1 {
		return createString(months, "month", end)
	}
	return createString(years, "year", end)
}

func createString(n float64, precision string, end string) string {
	n = math.Floor(n + 0.5)
	if n > 1 {
		return fmt.Sprintf("%d %ss %s", int(n), precision, end)
	}
	return fmt.Sprintf("%d %s %s", int(n), precision, end)
}
