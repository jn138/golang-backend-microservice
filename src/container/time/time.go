package time

import (
	"time"
)

type Time interface {
	Now() time.Time
	Since(t time.Time) time.Duration
}

type RealTime struct{}

func (t RealTime) Now() time.Time {
	return time.Now()
}

func (t RealTime) Since(s time.Time) time.Duration {
	return time.Since(s)
}
