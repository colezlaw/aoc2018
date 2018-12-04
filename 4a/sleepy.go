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

// Find the guard with the most moments sleep
func findSleepiestGuard(events []event) int {
	sleepTimes := make(map[int]int)
	start := time.Now()

	for _, e := range events {
		switch e.typ {
		case sleep:
			start = e.ts
		case wake:
			sleepTimes[e.id] += (int)(e.ts.Sub(start) / time.Minute)
		}
	}

	maxDur, maxID := 0, -1
	for k, v := range sleepTimes {
		if v > maxDur {
			maxDur = v
			maxID = k
		}
	}

	return maxID
}

func findSleepiestMinute(events []event, id int) int {
	start := time.Now()
	mins := make(map[int]int)
	for _, e := range events {
		if e.id != id {
			continue
		}
		switch e.typ {
		case sleep:
			start = e.ts
		case wake:
			for t := start; e.ts.After(t); t = t.Add(1 * time.Minute) {
				mins[t.Minute()]++
			}
		}
	}

	maxMin, maxVal := -1, -1
	for k, v := range mins {
		if v > maxVal {
			maxVal = v
			maxMin = k
		}
	}

	return maxMin
}
