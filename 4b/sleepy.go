package aoc

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"time"
)

type eventType uint32

const (
	beginShift eventType = iota
	sleep
	wake
	unknown
)

type event struct {
	ts  time.Time
	id  int
	typ eventType
}

// read the reader and sort.
func read(r io.Reader) []string {
	result := make([]string, 0)
	s := bufio.NewScanner(r)

	for s.Scan() {
		result = append(result, s.Text())
	}

	sort.Strings(result)
	return result
}

// Take a sorted list of event logs and turn it into a
// stream of events
func eventStream(s []string) []event {
	var id int
	events := make([]event, len(s))

	for i, l := range s {
		t, _ := time.Parse("2006-01-02 15:04", l[1:17])
		e := event{ts: t}

		switch l[19:24] {
		case "Guard":
			e.typ = beginShift
			fmt.Sscanf(l[26:], "%d", &id)
			e.id = id
		case "falls":
			e.typ = sleep
			e.id = id
		case "wakes":
			e.typ = wake
			e.id = id
		}

		events[i] = e
	}

	return events
}

// find the guard's minute that is sleepiest of all
func findSleepiestMinute(events []event) (guard, minute int) {
	start := time.Now()
	guardMins := make(map[int]map[int]int)
	for _, e := range events {
		if _, ok := guardMins[e.id]; !ok {
			guardMins[e.id] = make(map[int]int)
		}
		switch e.typ {
		case sleep:
			start = e.ts
		case wake:
			for t := start; e.ts.After(t); t = t.Add(1 * time.Minute) {
				guardMins[e.id][t.Minute()]++
			}
		}
	}

	maxMin, maxGuard, maxVal := -1, -1, -1
	for guard, v := range guardMins {
		for min, val := range v {
			if val > maxVal {
				maxMin = min
				maxGuard = guard
				maxVal = val
			}
		}
	}

	return maxGuard, maxMin
}
