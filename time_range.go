package timerange

import "time"

type TimeRange struct {
	Start time.Time
	End   time.Time
}

func (r TimeRange) SplitDaysAtTime(durationFromMidnight time.Duration) []TimeRange {
	ranges := []TimeRange{}
	start := r.Start
	day := 24 * 60 * 60 * time.Second

	d := BeginningOfDay(start).Add(durationFromMidnight)
	if !d.After(start) {
		d = d.Add(day)
	}
	for d.Before(r.End) {
		ranges = append(ranges, TimeRange{Start: start, End: d})
		start = d
		d = d.Add(day)
	}
	ranges = append(ranges, TimeRange{Start: start, End: r.End})
	return ranges
}

func BeginningOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func DurationFromMidnight(hhmmss string) (time.Duration, error) {
	t, err := time.Parse("2006-01-02 15:04:05", "0001-01-01 "+hhmmss)
	if err != nil {
		return 0, err
	}
	return t.Sub(time.Time{}), nil
}
