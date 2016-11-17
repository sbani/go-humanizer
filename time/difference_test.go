package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDifference(t *testing.T) {
	baseTime := time.Date(2016, 11, 3, 13, 0, 0, 0, time.UTC)
	assert.Equal(t,
		"just now",
		Difference(
			baseTime,
			baseTime,
		),
	)
	assert.Equal(t,
		"just now",
		Difference(
			baseTime,
			baseTime.Add(2*time.Nanosecond),
		),
	)
	assert.Equal(t,
		"5 seconds from now",
		Difference(
			baseTime,
			baseTime.Add(5*time.Second),
		),
	)
	assert.Equal(t,
		"1 minute ago",
		Difference(
			baseTime,
			baseTime.Add(-61*time.Second),
		),
	)
	assert.Equal(t,
		"15 minutes ago",
		Difference(
			baseTime,
			baseTime.Add(-(15*time.Minute+3*time.Nanosecond)),
		),
	)
	assert.Equal(t,
		"15 minutes from now",
		Difference(
			baseTime,
			baseTime.Add(14*time.Minute+50*time.Second),
		),
	)
	assert.Equal(t,
		"59 minutes ago",
		Difference(
			baseTime,
			baseTime.Add(-(59*time.Minute)),
		),
	)
	assert.Equal(t,
		"1 hour from now",
		Difference(
			baseTime,
			baseTime.Add(1*time.Hour+2*time.Minute),
		),
	)
	assert.Equal(t,
		"2 hours from now",
		Difference(
			baseTime,
			baseTime.Add(2*time.Hour+3*time.Minute),
		),
	)
	assert.Equal(t,
		"1 hour ago",
		Difference(
			baseTime,
			baseTime.Add(-(61*time.Minute)),
		),
	)
	assert.Equal(t,
		"1 day ago",
		Difference(
			baseTime,
			baseTime.Add(-(25*time.Hour)),
		),
	)
	assert.Equal(t,
		"2 days ago",
		Difference(
			baseTime,
			baseTime.Add(-(2*24*time.Hour+2*time.Hour)),
		),
	)
	assert.Equal(t,
		"1 day from now",
		Difference(
			baseTime,
			baseTime.Add(24*time.Hour),
		),
	)
	assert.Equal(t,
		"2 days from now",
		Difference(
			baseTime,
			baseTime.Add(47*time.Hour),
		),
	)
	assert.Equal(t,
		"1 week from now",
		Difference(
			baseTime,
			baseTime.Add(8*24*time.Hour),
		),
	)
	assert.Equal(t,
		"2 weeks from now",
		Difference(
			baseTime,
			baseTime.Add(14*24*time.Hour),
		),
	)
	assert.Equal(t,
		"1 month ago",
		Difference(
			baseTime,
			baseTime.Add(-31*24*time.Hour),
		),
	)
	assert.Equal(t,
		"5 months from now",
		Difference(
			baseTime,
			baseTime.Add(5*30*24*time.Hour-2*time.Hour),
		),
	)
	assert.Equal(t,
		"1 year from now",
		Difference(
			baseTime,
			baseTime.Add(366*24*time.Hour),
		),
	)
	assert.Equal(t,
		"2 years ago",
		Difference(
			baseTime,
			baseTime.Add(-(2*367*24*time.Hour+5*time.Hour)),
		),
	)
}

func TestCreateString(t *testing.T) {
	assert.Equal(t, "1 foo bar", createString(float64(1), "foo", "bar"))
	assert.Equal(t, "1 foo bar", createString(float64(1.002), "foo", "bar"))
	assert.Equal(t, "2 foos bar", createString(float64(2), "foo", "bar"))
	assert.Equal(t, "3 foos bar", createString(float64(2.9), "foo", "bar"))
	assert.Equal(t, "102 foos bar", createString(float64(101.5), "foo", "bar"))
	assert.Equal(t, "101 foos bar", createString(float64(101.4), "foo", "bar"))
}
