package time_test

import (
	Time "golang-backend-microservice/container/time"
	"testing"
	"time"
)

type MockTime struct{}

func (m MockTime) Now() time.Time {
	t, _ := time.Parse(Time.DATE_TIME_LAYOUT, "2023-01-01 04:30:18")
	return t
}

func (m MockTime) Since(s time.Time) time.Duration {
	t, _ := time.Parse(Time.DATE_TIME_LAYOUT, "2023-01-01 04:30:18")
	return t.Sub(s)
}

func TestTimeNow(t *testing.T) {
	expectedRealTime := time.Now()
	gotRealTime := Time.RealTime{}.Now()
	expectedMockTime, _ := time.Parse(Time.DATE_TIME_LAYOUT, "2023-01-01 04:30:18")
	gotMockTime := MockTime{}.Now()

	if expectedRealTime.Format(Time.DATE_TIME_LAYOUT) != gotRealTime.Format(Time.DATE_TIME_LAYOUT) {
		t.Errorf("Expected %v, got %v: incorrect real time", expectedRealTime.Format(Time.DATE_TIME_LAYOUT), gotRealTime.Format(Time.DATE_TIME_LAYOUT))
	}

	if expectedMockTime != gotMockTime {
		t.Errorf("Expected %v, got %v: incorrect mock time", expectedMockTime, gotMockTime)
	}
}

func TestTimeSince(t *testing.T) {
	var expectedMockTime time.Duration = 4*time.Hour + 30*time.Minute + 18*time.Second

	mockTimeSince, _ := time.Parse(Time.DATE_TIME_LAYOUT, "2023-01-01 00:00:00")
	gotMockTime := MockTime{}.Since(mockTimeSince)

	if expectedMockTime != gotMockTime {
		t.Errorf("Expected %v, got %v: incorrect mock time", expectedMockTime, gotMockTime)
	}
}
