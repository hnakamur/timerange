package timerange

import (
	"testing"
	"time"
)

func TestSplitDaysAtTimeOneRange(t *testing.T) {
	r := TimeRange{
		Start: time.Date(2015, 5, 31, 23, 59, 58, 0, time.Local),
		End:   time.Date(2015, 5, 31, 23, 59, 59, 0, time.Local),
	}
	d, err := DurationFromMidnight("00:01:05")
	if err != nil {
		t.Fatal(err)
	}
	got := r.SplitDaysAtTime(d)
	want := []TimeRange{
		TimeRange{
			Start: time.Date(2015, 5, 31, 23, 59, 58, 0, time.Local),
			End:   time.Date(2015, 5, 31, 23, 59, 59, 0, time.Local),
		},
	}
	checkTimeRanges(t, want, got)
}

func TestSplitDaysAtTimeSplitAtStart(t *testing.T) {
	r := TimeRange{
		Start: time.Date(2015, 5, 31, 23, 59, 58, 0, time.Local),
		End:   time.Date(2015, 6, 1, 23, 59, 59, 0, time.Local),
	}
	d, err := DurationFromMidnight("23:59:58")
	if err != nil {
		t.Fatal(err)
	}
	got := r.SplitDaysAtTime(d)
	want := []TimeRange{
		TimeRange{
			Start: time.Date(2015, 5, 31, 23, 59, 58, 0, time.Local),
			End:   time.Date(2015, 6, 1, 23, 59, 58, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 1, 23, 59, 58, 0, time.Local),
			End:   time.Date(2015, 6, 1, 23, 59, 59, 0, time.Local),
		},
	}
	checkTimeRanges(t, want, got)
}

func TestSplitDaysAtTimeSplitAtEnd(t *testing.T) {
	r := TimeRange{
		Start: time.Date(2015, 5, 31, 23, 59, 58, 0, time.Local),
		End:   time.Date(2015, 6, 1, 23, 59, 59, 0, time.Local),
	}
	d, err := DurationFromMidnight("23:59:59")
	if err != nil {
		t.Fatal(err)
	}
	got := r.SplitDaysAtTime(d)
	want := []TimeRange{
		TimeRange{
			Start: time.Date(2015, 5, 31, 23, 59, 58, 0, time.Local),
			End:   time.Date(2015, 5, 31, 23, 59, 59, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 5, 31, 23, 59, 59, 0, time.Local),
			End:   time.Date(2015, 6, 1, 23, 59, 59, 0, time.Local),
		},
	}
	checkTimeRanges(t, want, got)
}

func TestSplitDaysAtTimeMonth(t *testing.T) {
	r := TimeRange{
		Start: time.Date(2015, 5, 31, 23, 59, 58, 0, time.Local),
		End:   time.Date(2015, 7, 1, 0, 2, 15, 0, time.Local),
	}
	d, err := DurationFromMidnight("00:01:05")
	if err != nil {
		t.Fatal(err)
	}
	got := r.SplitDaysAtTime(d)
	want := []TimeRange{
		TimeRange{
			Start: time.Date(2015, 5, 31, 23, 59, 58, 0, time.Local),
			End:   time.Date(2015, 6, 1, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 1, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 2, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 2, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 3, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 3, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 4, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 4, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 5, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 5, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 6, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 6, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 7, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 7, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 8, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 8, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 9, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 9, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 10, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 10, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 11, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 11, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 12, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 12, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 13, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 13, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 14, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 14, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 15, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 15, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 16, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 16, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 17, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 17, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 18, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 18, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 19, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 19, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 20, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 20, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 21, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 21, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 22, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 22, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 23, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 23, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 24, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 24, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 25, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 25, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 26, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 26, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 27, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 27, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 28, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 28, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 29, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 29, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 6, 30, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 6, 30, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 7, 1, 0, 1, 5, 0, time.Local),
		},
		TimeRange{
			Start: time.Date(2015, 7, 1, 0, 1, 5, 0, time.Local),
			End:   time.Date(2015, 7, 1, 0, 2, 15, 0, time.Local),
		},
	}
	checkTimeRanges(t, want, got)
}

func checkTimeRanges(t *testing.T, want []TimeRange, got []TimeRange) {
	if len(want) != len(got) {
		t.Errorf("length mismatch: want=%d, got=%d", len(want), len(got))
	}
	for i, w := range want {
		g := got[i]
		if w.Start != g.Start {
			t.Errorf("start time mismatch: i=%d, want=%v, got=%v", i, w.Start, g.Start)
		}
		if w.End != g.End {
			t.Errorf("end time mismatch: i=%d, want=%v, got=%v", i, w.End, g.End)
		}
	}
}
