package aoc

import (
	"os"
	"path"
	"strings"
	"testing"
	"time"
)

var input = `[1518-11-05 00:55] wakes up
[1518-11-05 00:45] falls asleep
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-04 00:46] wakes up
[1518-11-04 00:36] falls asleep
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-03 00:29] wakes up
[1518-11-03 00:24] falls asleep
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-02 00:50] wakes up
[1518-11-02 00:40] falls asleep
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-01 00:55] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:05] falls asleep
[1518-11-01 00:00] Guard #10 begins shift`

// First, verify we sort correctly
func TestRead(t *testing.T) {
	tt := []string{
		"[1518-11-01 00:00]",
		"[1518-11-01 00:05]",
		"[1518-11-01 00:25]",
		"[1518-11-01 00:30]",
		"[1518-11-01 00:55]",
		"[1518-11-01 23:58]",
		"[1518-11-02 00:40]",
		"[1518-11-02 00:50]",
		"[1518-11-03 00:05]",
		"[1518-11-03 00:24]",
		"[1518-11-03 00:29]",
		"[1518-11-04 00:02]",
		"[1518-11-04 00:36]",
		"[1518-11-04 00:46]",
		"[1518-11-05 00:03]",
		"[1518-11-05 00:45]",
		"[1518-11-05 00:55]",
	}
	result := read(strings.NewReader(input))
	if len(result) != len(tt) {
		t.Fatalf("expected len of %d, got %d", len(tt), len(result))
	}

	for i, s := range tt {
		if result[i][:len(tt)+1] != s {
			t.Errorf("Expected %s, got %s", s, result[i][:len(tt)+1])
		}
	}
}

// Given a sorted array of event log entries, test to ensure
// we get back an array of the correct events
func TestEventStream(t *testing.T) {
	logs := []string{
		"[1518-11-01 00:00] Guard #10 begins shift",
		"[1518-11-01 00:05] falls asleep",
		"[1518-11-01 00:25] wakes up",
		"[1518-11-01 00:30] falls asleep",
		"[1518-11-01 00:55] wakes up",
		"[1518-11-01 23:58] Guard #99 begins shift",
		"[1518-11-02 00:40] falls asleep",
		"[1518-11-02 00:50] wakes up",
		"[1518-11-03 00:05] Guard #10 begins shift",
		"[1518-11-03 00:24] falls asleep",
		"[1518-11-03 00:29] wakes up",
		"[1518-11-04 00:02] Guard #99 begins shift",
		"[1518-11-04 00:36] falls asleep",
		"[1518-11-04 00:46] wakes up",
		"[1518-11-05 00:03] Guard #99 begins shift",
		"[1518-11-05 00:45] falls asleep",
		"[1518-11-05 00:55] wakes up",
	}
	tt := []event{
		event{ts: time.Date(1518, 11, 1, 0, 0, 0, 0, time.UTC), typ: beginShift, id: 10},
		event{ts: time.Date(1518, 11, 1, 0, 5, 0, 0, time.UTC), typ: sleep, id: 10},
		event{ts: time.Date(1518, 11, 1, 0, 25, 0, 0, time.UTC), typ: wake, id: 10},
		event{ts: time.Date(1518, 11, 1, 0, 30, 0, 0, time.UTC), typ: sleep, id: 10},
		event{ts: time.Date(1518, 11, 1, 0, 55, 0, 0, time.UTC), typ: wake, id: 10},
		event{ts: time.Date(1518, 11, 1, 23, 58, 0, 0, time.UTC), typ: beginShift, id: 99},
		event{ts: time.Date(1518, 11, 2, 0, 40, 0, 0, time.UTC), typ: sleep, id: 99},
		event{ts: time.Date(1518, 11, 2, 0, 50, 0, 0, time.UTC), typ: wake, id: 99},
		event{ts: time.Date(1518, 11, 3, 0, 5, 0, 0, time.UTC), typ: beginShift, id: 10},
		event{ts: time.Date(1518, 11, 3, 0, 24, 0, 0, time.UTC), typ: sleep, id: 10},
		event{ts: time.Date(1518, 11, 3, 0, 29, 0, 0, time.UTC), typ: wake, id: 10},
		event{ts: time.Date(1518, 11, 4, 0, 2, 0, 0, time.UTC), typ: beginShift, id: 99},
		event{ts: time.Date(1518, 11, 4, 0, 36, 0, 0, time.UTC), typ: sleep, id: 99},
		event{ts: time.Date(1518, 11, 4, 0, 46, 0, 0, time.UTC), typ: wake, id: 99},
		event{ts: time.Date(1518, 11, 5, 0, 03, 0, 0, time.UTC), typ: beginShift, id: 99},
		event{ts: time.Date(1518, 11, 5, 0, 45, 0, 0, time.UTC), typ: sleep, id: 99},
		event{ts: time.Date(1518, 11, 5, 0, 55, 0, 0, time.UTC), typ: wake, id: 99},
	}

	events := eventStream(logs)
	for i, tc := range tt {
		if tc != events[i] {
			t.Errorf("expected %v got %v", tc, events[i])
		}
	}
}

func TestFindSleepiestGuard(t *testing.T) {
	tt := []event{
		event{ts: time.Date(1518, 11, 1, 0, 0, 0, 0, time.UTC), typ: beginShift, id: 10},
		event{ts: time.Date(1518, 11, 1, 0, 5, 0, 0, time.UTC), typ: sleep, id: 10},
		event{ts: time.Date(1518, 11, 1, 0, 25, 0, 0, time.UTC), typ: wake, id: 10},
		event{ts: time.Date(1518, 11, 1, 0, 30, 0, 0, time.UTC), typ: sleep, id: 10},
		event{ts: time.Date(1518, 11, 1, 0, 55, 0, 0, time.UTC), typ: wake, id: 10},
		event{ts: time.Date(1518, 11, 1, 23, 58, 0, 0, time.UTC), typ: beginShift, id: 99},
		event{ts: time.Date(1518, 11, 2, 0, 40, 0, 0, time.UTC), typ: sleep, id: 99},
		event{ts: time.Date(1518, 11, 2, 0, 50, 0, 0, time.UTC), typ: wake, id: 99},
		event{ts: time.Date(1518, 11, 3, 0, 5, 0, 0, time.UTC), typ: beginShift, id: 10},
		event{ts: time.Date(1518, 11, 3, 0, 24, 0, 0, time.UTC), typ: sleep, id: 10},
		event{ts: time.Date(1518, 11, 3, 0, 29, 0, 0, time.UTC), typ: wake, id: 10},
		event{ts: time.Date(1518, 11, 4, 0, 2, 0, 0, time.UTC), typ: beginShift, id: 99},
		event{ts: time.Date(1518, 11, 4, 0, 36, 0, 0, time.UTC), typ: sleep, id: 99},
		event{ts: time.Date(1518, 11, 4, 0, 46, 0, 0, time.UTC), typ: wake, id: 99},
		event{ts: time.Date(1518, 11, 5, 0, 03, 0, 0, time.UTC), typ: beginShift, id: 99},
		event{ts: time.Date(1518, 11, 5, 0, 45, 0, 0, time.UTC), typ: sleep, id: 99},
		event{ts: time.Date(1518, 11, 5, 0, 55, 0, 0, time.UTC), typ: wake, id: 99},
	}
	result := findSleepiestGuard(tt)
	if result != 10 {
		t.Errorf("expected 10, got %d", result)
	}
}

func TestFindSleepiestMinute(t *testing.T) {
	tt := []event{
		event{ts: time.Date(1518, 11, 1, 0, 0, 0, 0, time.UTC), typ: beginShift, id: 10},
		event{ts: time.Date(1518, 11, 1, 0, 5, 0, 0, time.UTC), typ: sleep, id: 10},
		event{ts: time.Date(1518, 11, 1, 0, 25, 0, 0, time.UTC), typ: wake, id: 10},
		event{ts: time.Date(1518, 11, 1, 0, 30, 0, 0, time.UTC), typ: sleep, id: 10},
		event{ts: time.Date(1518, 11, 1, 0, 55, 0, 0, time.UTC), typ: wake, id: 10},
		event{ts: time.Date(1518, 11, 1, 23, 58, 0, 0, time.UTC), typ: beginShift, id: 99},
		event{ts: time.Date(1518, 11, 2, 0, 40, 0, 0, time.UTC), typ: sleep, id: 99},
		event{ts: time.Date(1518, 11, 2, 0, 50, 0, 0, time.UTC), typ: wake, id: 99},
		event{ts: time.Date(1518, 11, 3, 0, 5, 0, 0, time.UTC), typ: beginShift, id: 10},
		event{ts: time.Date(1518, 11, 3, 0, 24, 0, 0, time.UTC), typ: sleep, id: 10},
		event{ts: time.Date(1518, 11, 3, 0, 29, 0, 0, time.UTC), typ: wake, id: 10},
		event{ts: time.Date(1518, 11, 4, 0, 2, 0, 0, time.UTC), typ: beginShift, id: 99},
		event{ts: time.Date(1518, 11, 4, 0, 36, 0, 0, time.UTC), typ: sleep, id: 99},
		event{ts: time.Date(1518, 11, 4, 0, 46, 0, 0, time.UTC), typ: wake, id: 99},
		event{ts: time.Date(1518, 11, 5, 0, 03, 0, 0, time.UTC), typ: beginShift, id: 99},
		event{ts: time.Date(1518, 11, 5, 0, 45, 0, 0, time.UTC), typ: sleep, id: 99},
		event{ts: time.Date(1518, 11, 5, 0, 55, 0, 0, time.UTC), typ: wake, id: 99},
	}
	result := findSleepiestMinute(tt, 10)
	if result != 24 {
		t.Errorf("expected 24, got %d", result)
	}
}

func TestExam(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	f, err := os.Open(path.Join("testdata", "input.txt"))
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	defer f.Close()

	logs := read(f)
	events := eventStream(logs)
	guard := findSleepiestGuard(events)
	minute := findSleepiestMinute(events, guard)
	t.Logf("The answer is %d", guard*minute)
}
